import React, { useState } from 'react';
import { X, Book, DollarSign, Hash } from 'lucide-react';
import { api } from '../api';

interface AddBookModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSuccess: () => void;
}

export function AddBookModal({ isOpen, onClose, onSuccess }: AddBookModalProps) {
  const [title, setTitle] = useState('');
  const [price, setPrice] = useState('');
  const [isbn, setIsbn] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  if (!isOpen) return null;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setIsLoading(true);

    try {
      await api.createBook({
        title,
        price: Number(price),
        isbn
      });
      onSuccess();
      onClose();
      setTitle('');
      setPrice('');
      setIsbn('');
    } catch (err: any) {
      setError(err.error?.message || 'Failed to add book.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/20 backdrop-blur-sm">
      <div className="bg-white rounded-3xl shadow-xl w-full max-w-md overflow-hidden border border-[var(--color-latte-100)] relative">
        <div className="absolute top-0 left-0 right-0 h-2 bg-[var(--color-sage-500)]"></div>
        
        <div className="p-6 flex justify-between items-center border-b border-[var(--color-latte-50)]">
          <h2 className="text-xl font-bold text-[var(--color-walnut-900)]">Add New Book</h2>
          <button onClick={onClose} className="text-[var(--color-walnut-800)] opacity-50 hover:opacity-100 transition-opacity">
            <X size={24} />
          </button>
        </div>

        <div className="p-6">
          {error && (
            <div className="mb-6 p-4 bg-red-50 text-red-600 rounded-2xl text-sm border border-red-100">
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-5">
            <div>
              <label className="block text-sm font-medium text-[var(--color-walnut-800)] mb-2">Title</label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-[var(--color-walnut-800)] opacity-50">
                  <Book size={18} />
                </div>
                <input type="text" required value={title} onChange={e => setTitle(e.target.value)} className="block w-full pl-11 pr-4 py-3 bg-[var(--color-latte-50)] border-transparent rounded-2xl text-[var(--color-walnut-900)] focus:bg-white focus:border-[var(--color-sage-500)] focus:ring-2 focus:ring-[var(--color-sage-500)] focus:ring-opacity-20 transition-all outline-none" placeholder="Domain Driven Design" />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-[var(--color-walnut-800)] mb-2">Price (¥)</label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-[var(--color-walnut-800)] opacity-50">
                  <DollarSign size={18} />
                </div>
                <input type="number" min="0" required value={price} onChange={e => setPrice(e.target.value)} className="block w-full pl-11 pr-4 py-3 bg-[var(--color-latte-50)] border-transparent rounded-2xl text-[var(--color-walnut-900)] focus:bg-white focus:border-[var(--color-sage-500)] focus:ring-2 focus:ring-[var(--color-sage-500)] focus:ring-opacity-20 transition-all outline-none" placeholder="4200" />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-[var(--color-walnut-800)] mb-2">ISBN</label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-[var(--color-walnut-800)] opacity-50">
                  <Hash size={18} />
                </div>
                <input type="text" required value={isbn} onChange={e => setIsbn(e.target.value)} className="block w-full pl-11 pr-4 py-3 bg-[var(--color-latte-50)] border-transparent rounded-2xl text-[var(--color-walnut-900)] focus:bg-white focus:border-[var(--color-sage-500)] focus:ring-2 focus:ring-[var(--color-sage-500)] focus:ring-opacity-20 transition-all outline-none" placeholder="978-4798121963" />
              </div>
            </div>

            <div className="pt-2 flex gap-3">
              <button type="button" onClick={onClose} className="flex-1 py-3 px-4 border border-[var(--color-latte-100)] rounded-2xl text-sm font-medium text-[var(--color-walnut-800)] hover:bg-[var(--color-latte-50)] transition-all">Cancel</button>
              <button type="submit" disabled={isLoading} className="flex-1 py-3 px-4 border border-transparent rounded-2xl shadow-sm text-sm font-medium text-white bg-[var(--color-sage-500)] hover:bg-[var(--color-sage-600)] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[var(--color-sage-500)] transition-all disabled:opacity-70">
                {isLoading ? 'Adding...' : 'Add Book'}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}
