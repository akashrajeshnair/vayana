"use client";

import type { ChangeEvent, FormEvent } from "react";
import { useEffect, useState } from "react";
import { useAuth } from "@/context/AuthContext";

type Book = {
  id: number;
  title: string;
  author: string;
  genre: string;
  coverImageUrl?: string | null;
  description?: string | null;
};

type ShelfStatus = "WANT_TO_READ" | "READING" | "FINISHED";

const statusOptions: { label: string; value: ShelfStatus }[] = [
  { label: "Want to Read", value: "WANT_TO_READ" },
  { label: "Reading", value: "READING" },
  { label: "Finished", value: "FINISHED" }
];

export default function BooksPage() {
  const { token } = useAuth();
  const [books, setBooks] = useState<Book[]>([]);
  const [query, setQuery] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [openDropdownId, setOpenDropdownId] = useState<number | null>(null);

  const fetchBooks = async (searchQuery?: string) => {
    setLoading(true);
    setError(null);
    try {
      const baseUrl = process.env.NEXT_PUBLIC_API_URL;
      const url = searchQuery
        ? `${baseUrl}/api/books/search?query=${encodeURIComponent(searchQuery)}`
        : `${baseUrl}/api/books`;

      const response = await fetch(url, {
        headers: {
          Authorization: token ? `Bearer ${token}` : ""
        }
      });

      if (!response.ok) {
        setError("Unable to load books.");
        return;
      }

      const data: Book[] = await response.json();
      setBooks(data);
    } catch (err) {
      setError("Unable to load books.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  const handleSearch = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    fetchBooks(query.trim());
  };

  const handleAddToShelf = async (bookId: number, status: ShelfStatus) => {
    try {
      const baseUrl = process.env.NEXT_PUBLIC_API_URL;
      const response = await fetch(`${baseUrl}/api/shelf`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: token ? `Bearer ${token}` : ""
        },
        body: JSON.stringify({ bookId, status })
      });

      if (!response.ok) {
        setError("Unable to add book to shelf.");
      }
    } catch (err) {
      setError("Unable to add book to shelf.");
    } finally {
      setOpenDropdownId(null);
    }
  };

  return (
    <main className="mx-auto max-w-5xl p-6">
      <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <h1 className="text-2xl font-semibold">Books</h1>
          <p className="mt-1 text-slate-600">Browse and search books in the catalog.</p>
        </div>
        <form className="flex w-full gap-2 sm:w-auto" onSubmit={handleSearch}>
          <input
            className="w-full rounded border border-slate-300 px-3 py-2 text-sm"
            placeholder="Search by title or author"
            value={query}
            onChange={(event: ChangeEvent<HTMLInputElement>) => setQuery(event.target.value)}
          />
          <button
            className="rounded bg-slate-900 px-4 py-2 text-sm text-white"
            type="submit"
          >
            Search
          </button>
        </form>
      </div>

      {error ? <p className="mt-4 text-sm text-red-600">{error}</p> : null}
      {loading ? <p className="mt-4 text-sm text-slate-600">Loading books...</p> : null}

      <div className="mt-6 grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        {books.map((book) => (
          <div key={book.id} className="rounded border border-slate-200 bg-white p-4 shadow-sm">
            <div className="flex gap-4">
              <div className="h-20 w-14 overflow-hidden rounded bg-slate-200">
                {book.coverImageUrl ? (
                  <img
                    src={book.coverImageUrl}
                    alt={book.title}
                    className="h-full w-full object-cover"
                  />
                ) : (
                  <div className="flex h-full w-full items-center justify-center text-xs text-slate-500">
                    No cover
                  </div>
                )}
              </div>
              <div className="flex-1">
                <h2 className="text-lg font-semibold">{book.title}</h2>
                <p className="text-sm text-slate-600">{book.author}</p>
                <p className="mt-1 text-xs uppercase text-slate-500">{book.genre}</p>
              </div>
            </div>
            <div className="mt-4">
              <button
                className="rounded border border-slate-300 px-3 py-2 text-sm"
                type="button"
                onClick={() =>
                  setOpenDropdownId(openDropdownId === book.id ? null : book.id)
                }
              >
                Add to Shelf
              </button>
              {openDropdownId === book.id ? (
                <div className="mt-2 rounded border border-slate-200 bg-white p-2 shadow">
                  {statusOptions.map((option) => (
                    <button
                      key={option.value}
                      className="block w-full rounded px-2 py-1 text-left text-sm hover:bg-slate-100"
                      type="button"
                      onClick={() => handleAddToShelf(book.id, option.value)}
                    >
                      {option.label}
                    </button>
                  ))}
                </div>
              ) : null}
            </div>
          </div>
        ))}
      </div>
    </main>
  );
}
