import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Coffee, Lock, Mail, User } from 'lucide-react';
import { api } from '../api';

export function SignUp() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setIsLoading(true);

    try {
      await api.signup({ name, email, password });
      navigate('/login');
    } catch (err: any) {
      setError(err.error?.message || 'Failed to sign up. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-[var(--color-latte-50)] p-4">
      <div className="w-full max-w-md bg-white rounded-3xl shadow-sm overflow-hidden border border-[var(--color-latte-100)] relative">
        <div className="absolute top-0 left-0 right-0 h-2 bg-[var(--color-walnut-800)]"></div>
        
        <div className="p-10">
          <div className="flex justify-center mb-8">
            <div className="bg-[var(--color-sage-100)] p-4 rounded-2xl text-[var(--color-sage-600)] shadow-sm">
              <Coffee size={32} strokeWidth={2.5} />
            </div>
          </div>
          
          <h2 className="text-3xl font-bold text-center text-[var(--color-walnut-900)] mb-2 tracking-tight">
            Join Librigo
          </h2>
          <p className="text-center text-[var(--color-walnut-800)] opacity-70 mb-8">
            Register to start adding to our collection.
          </p>

          {error && (
            <div className="mb-6 p-4 bg-red-50 text-red-600 rounded-2xl text-sm border border-red-100">
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <label className="block text-sm font-medium text-[var(--color-walnut-800)] mb-2">
                Full Name
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-[var(--color-walnut-800)] opacity-50">
                  <User size={20} />
                </div>
                <input
                  type="text"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  required
                  className="block w-full pl-11 pr-4 py-3 bg-[var(--color-latte-50)] border-transparent rounded-2xl text-[var(--color-walnut-900)] focus:bg-white focus:border-[var(--color-sage-500)] focus:ring-2 focus:ring-[var(--color-sage-500)] focus:ring-opacity-20 transition-all outline-none"
                  placeholder="Taro Yamada"
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-[var(--color-walnut-800)] mb-2">
                Email Address
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-[var(--color-walnut-800)] opacity-50">
                  <Mail size={20} />
                </div>
                <input
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                  className="block w-full pl-11 pr-4 py-3 bg-[var(--color-latte-50)] border-transparent rounded-2xl text-[var(--color-walnut-900)] focus:bg-white focus:border-[var(--color-sage-500)] focus:ring-2 focus:ring-[var(--color-sage-500)] focus:ring-opacity-20 transition-all outline-none"
                  placeholder="barista@librigo.cafe"
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-[var(--color-walnut-800)] mb-2">
                Password
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none text-[var(--color-walnut-800)] opacity-50">
                  <Lock size={20} />
                </div>
                <input
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                  minLength={8}
                  className="block w-full pl-11 pr-4 py-3 bg-[var(--color-latte-50)] border-transparent rounded-2xl text-[var(--color-walnut-900)] focus:bg-white focus:border-[var(--color-sage-500)] focus:ring-2 focus:ring-[var(--color-sage-500)] focus:ring-opacity-20 transition-all outline-none"
                  placeholder="••••••••"
                />
              </div>
            </div>

            <button
              type="submit"
              disabled={isLoading}
              className="w-full flex justify-center py-4 px-4 border border-transparent rounded-2xl shadow-sm text-base font-medium text-white bg-[var(--color-sage-500)] hover:bg-[var(--color-sage-600)] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[var(--color-sage-500)] transition-all disabled:opacity-70 disabled:cursor-not-allowed"
            >
              {isLoading ? 'Registering...' : 'Sign Up'}
            </button>
          </form>

          <p className="text-center text-sm text-[var(--color-walnut-800)] mt-8">
            Already have an account?{' '}
            <button
              onClick={() => navigate('/login')}
              className="text-[var(--color-sage-600)] font-medium hover:underline"
            >
              Sign in here
            </button>
          </p>
        </div>
      </div>
    </div>
  );
}
