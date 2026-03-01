import React, { ReactNode } from 'react';
import { BookOpen, Library, LogOut, LogIn, UserPlus, Coffee } from 'lucide-react';
import { Link, useLocation, useNavigate } from 'react-router-dom';

interface LayoutProps {
  children: ReactNode;
}

export function Layout({ children }: LayoutProps) {
  const location = useLocation();
  const navigate = useNavigate();
  const isAuthenticated = !!localStorage.getItem('librigo_token');

  const handleLogout = () => {
    localStorage.removeItem('librigo_token');
    navigate('/login');
  };

  const navItems = [
    { path: '/', icon: Library, label: 'Dashboard' },
    { path: '/books', icon: BookOpen, label: 'Books' },
  ];

  return (
    <div className="flex h-screen bg-[var(--color-latte-50)] overflow-hidden">
      {/* Sidebar */}
      <aside className="w-64 bg-white border-r border-[var(--color-latte-100)] flex flex-col justify-between shadow-sm z-10 relative">
        {/* Wood-grain accent top bar */}
        <div className="absolute top-0 left-0 right-0 h-1 bg-[var(--color-walnut-800)]"></div>
        
        <div>
          <div className="p-8 flex items-center gap-3">
            <div className="bg-[var(--color-sage-100)] p-2 rounded-xl text-[var(--color-sage-600)]">
              <Coffee size={24} strokeWidth={2.5} />
            </div>
            <h1 className="text-2xl font-bold text-[var(--color-walnut-900)] tracking-tight">Librigo</h1>
          </div>

          <nav className="px-4 space-y-2">
            {navItems.map((item) => {
              const isActive = location.pathname === item.path;
              return (
                <Link
                  key={item.path}
                  to={item.path}
                  className={`flex items-center gap-3 px-4 py-3 rounded-2xl transition-all duration-200 ${
                    isActive
                      ? 'bg-[var(--color-sage-500)] text-white shadow-sm'
                      : 'text-[var(--color-walnut-800)] hover:bg-[var(--color-latte-100)]'
                  }`}
                >
                  <item.icon size={20} strokeWidth={isActive ? 2.5 : 2} />
                  <span className="font-medium">{item.label}</span>
                </Link>
              );
            })}
          </nav>
        </div>

        <div className="p-4 space-y-2">
          {isAuthenticated ? (
            <button
              onClick={handleLogout}
              className="flex items-center gap-3 px-4 py-3 w-full rounded-2xl text-[var(--color-walnut-800)] hover:bg-[var(--color-latte-100)] transition-colors"
            >
              <LogOut size={20} />
              <span className="font-medium">Logout</span>
            </button>
          ) : (
            <>
              <Link
                to="/login"
                className="flex items-center gap-3 px-4 py-3 w-full rounded-2xl text-[var(--color-walnut-800)] hover:bg-[var(--color-latte-100)] transition-colors"
              >
                <LogIn size={20} />
                <span className="font-medium">Login</span>
              </Link>
              <Link
                to="/signup"
                className="flex items-center gap-3 px-4 py-3 w-full rounded-2xl text-white bg-[var(--color-sage-500)] hover:bg-[var(--color-sage-600)] transition-colors shadow-sm"
              >
                <UserPlus size={20} />
                <span className="font-medium">Sign Up</span>
              </Link>
            </>
          )}
        </div>
      </aside>

      {/* Main Content */}
      <main className="flex-1 overflow-y-auto p-8 md:p-12">
        <div className="max-w-6xl mx-auto">
          {children}
        </div>
      </main>
    </div>
  );
}
