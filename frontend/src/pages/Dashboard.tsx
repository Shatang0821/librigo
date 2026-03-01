import React, { useEffect, useState } from 'react';
import { BookOpen, Plus, Search, Coffee } from 'lucide-react';
import { api } from '../api';
import type { BookResponse } from '../api';
import { AddBookModal } from '../components/AddBookModal';

export function Dashboard() {
  const [books, setBooks] = useState<BookResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState('');
  const [searchQuery, setSearchQuery] = useState('');
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  
  const isAuthenticated = !!localStorage.getItem('librigo_token');

  const fetchBooks = async () => {
    setIsLoading(true);
    try {
      const data = await api.getBooks();
      setBooks(data);
    } catch (err: any) {
      setError(err.error?.message || 'Failed to load books.');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  const filteredBooks = books.filter(book => 
    book.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
    book.isbn.includes(searchQuery)
  );

  return (
    <div className="space-y-8">
      {/* Header Section */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 bg-white p-8 rounded-3xl shadow-sm border border-[var(--color-latte-100)] relative overflow-hidden">
        <div className="absolute top-0 left-0 bottom-0 w-2 bg-[var(--color-sage-500)]"></div>
        <div>
          <h1 className="text-4xl font-bold text-[var(--color-walnut-900)] tracking-tight mb-2">
            Library Menu
          </h1>
          <p className="text-[var(--color-walnut-800)] opacity-70 flex items-center gap-2">
            <Coffee size={16} />
            Discover our freshly brewed collection of books.
          </p>
        </div>
        {isAuthenticated && (
          <button 
            onClick={() => setIsAddModalOpen(true)}
            className="flex items-center gap-2 bg-[var(--color-walnut-800)] hover:bg-[var(--color-walnut-900)] text-white px-6 py-3 rounded-2xl shadow-sm transition-all font-medium"
          >
            <Plus size={20} />
            Add New Book
          </button>
        )}
      </div>

      {/* Search and Filter */}
      <div className="flex gap-4">
        <div className="relative flex-1 max-w-md">
          <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-[var(--color-walnut-800)] opacity-50">
            <Search size={20} />
          </div>
          <input
            type="text"
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            placeholder="Search by title or ISBN..."
            className="block w-full pl-11 pr-4 py-3 bg-white border border-[var(--color-latte-100)] rounded-2xl text-[var(--color-walnut-900)] focus:border-[var(--color-sage-500)] focus:ring-2 focus:ring-[var(--color-sage-500)] focus:ring-opacity-20 transition-all outline-none shadow-sm"
          />
        </div>
      </div>

      {/* Error State */}
      {error && (
        <div className="p-6 bg-red-50 text-red-600 rounded-2xl border border-red-100">
          {error}
        </div>
      )}

      {/* Loading State */}
      {isLoading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {[1, 2, 3, 4, 5, 6].map((i) => (
            <div key={i} className="bg-white p-6 rounded-3xl shadow-sm border border-[var(--color-latte-100)] animate-pulse h-48">
              <div className="h-6 bg-[var(--color-latte-100)] rounded w-3/4 mb-4"></div>
              <div className="h-4 bg-[var(--color-latte-100)] rounded w-1/2 mb-8"></div>
              <div className="h-4 bg-[var(--color-latte-100)] rounded w-1/4"></div>
            </div>
          ))}
        </div>
      ) : (
        /* Book Grid */
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {filteredBooks.length === 0 ? (
            <div className="col-span-full bg-white p-12 rounded-3xl shadow-sm border border-[var(--color-latte-100)] text-center">
              <div className="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-[var(--color-latte-50)] text-[var(--color-walnut-800)] opacity-50 mb-4">
                <BookOpen size={32} />
              </div>
              <h3 className="text-xl font-semibold text-[var(--color-walnut-900)] mb-2">No books found</h3>
              <p className="text-[var(--color-walnut-800)] opacity-70">No books found matching your search.</p>
            </div>
          ) : (
            filteredBooks.map((book) => (
              <div
                key={book.id}
                className="group bg-white p-6 rounded-3xl shadow-sm border border-[var(--color-latte-100)] hover:shadow-md hover:border-[var(--color-sage-500)] transition-all duration-300 relative overflow-hidden flex flex-col justify-between h-full"
              >
                {/* Decorative wood accent on hover */}
                <div className="absolute top-0 left-0 right-0 h-1 bg-[var(--color-walnut-800)] transform -translate-y-full group-hover:translate-y-0 transition-transform duration-300"></div>
                
                <div>
                  <div className="flex justify-between items-start mb-4">
                    <div className="bg-[var(--color-latte-50)] p-3 rounded-2xl text-[var(--color-sage-600)] group-hover:bg-[var(--color-sage-100)] transition-colors">
                      <BookOpen size={24} strokeWidth={2} />
                    </div>
                    <span className="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-[var(--color-sage-100)] text-[var(--color-sage-600)]">
                      Available
                    </span>
                  </div>
                  
                  <h3 className="text-xl font-bold text-[var(--color-walnut-900)] mb-1 line-clamp-2 leading-tight">
                    {book.title}
                  </h3>
                  <p className="text-sm text-[var(--color-walnut-800)] opacity-60 font-mono mb-6">
                    ISBN: {book.isbn}
                  </p>
                </div>

                <div className="flex items-end justify-between mt-4 pt-4 border-t border-[var(--color-latte-50)]">
                  <div>
                    <p className="text-xs text-[var(--color-walnut-800)] opacity-60 uppercase tracking-wider font-semibold mb-1">
                      Price
                    </p>
                    <p className="text-2xl font-bold text-[var(--color-sage-600)]">
                      ¥{book.price.toLocaleString()}
                    </p>
                  </div>
                </div>
              </div>
            ))
          )}
        </div>
      )}

      <AddBookModal 
        isOpen={isAddModalOpen} 
        onClose={() => setIsAddModalOpen(false)} 
        onSuccess={fetchBooks} 
      />
    </div>
  );
}
