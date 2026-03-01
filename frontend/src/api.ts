export type Book = {
  id: string;
  title: string;
  price: number;
  isbn: string;
};

export type User = {
  id: string;
  name: string;
  email: string;
  role?: 'admin' | 'member';
};

export type ApiError = {
  error: {
    code: string;
    message: string;
  };
};

export type SignUpRequest = {
  name: string;
  email: string;
  password: string;
};

export type SignInRequest = {
  email: string;
  password: string;
};

export type CreateBookRequest = {
  title: string;
  price: number;
  isbn: string;
};

export type SignUpResponse = {
  id: string;
  name: string;
  email: string;
};

export type SignInResponse = {
  token: string;
  user: {
    id: string;
    name: string;
    email: string;
  };
};

export type CreateBookResponse = {
  id: string;
  title: string;
};

export type BookResponse = Book;

const BASE_URL = 'http://localhost:8080';

async function fetchApi<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const token = localStorage.getItem('librigo_token');
  const headers: HeadersInit = {
    'Content-Type': 'application/json',
    ...options.headers,
  };

  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  const response = await fetch(`${BASE_URL}${endpoint}`, {
    ...options,
    headers,
  });

  if (!response.ok) {
    let errorData: ApiError;
    try {
      errorData = await response.json();
    } catch {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    throw errorData;
  }

  // Handle empty responses (like 201 Created with no body, though backend says it returns JSON)
  const text = await response.text();
  return text ? JSON.parse(text) : null;
}

export const api = {
  signup: (data: SignUpRequest) => fetchApi<SignUpResponse>('/signup', {
    method: 'POST',
    body: JSON.stringify(data),
  }),
  signin: (data: SignInRequest) => fetchApi<SignInResponse>('/signin', {
    method: 'POST',
    body: JSON.stringify(data),
  }),
  getBooks: () => fetchApi<BookResponse[]>('/books', {
    method: 'GET',
  }),
  getBook: (id: string) => fetchApi<BookResponse>(`/books/${id}`, {
    method: 'GET',
  }),
  createBook: (data: CreateBookRequest) => fetchApi<CreateBookResponse>('/books', {
    method: 'POST',
    body: JSON.stringify(data),
  }),
};
