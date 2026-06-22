"use client";

import { useEffect, useMemo, useState } from "react";
import { useAuth } from "@/context/AuthContext";

type ShelfRecord = {
  id: number;
  bookId: number;
  status: "WANT_TO_READ" | "READING" | "FINISHED";
  rating?: number | null;
  review?: string | null;
  startDate?: string | null;
  finishDate?: string | null;
};

type BookSummary = {
  id: number;
  title: string;
  author: string;
};

const statusLabels: Record<ShelfRecord["status"], string> = {
  WANT_TO_READ: "Want to Read",
  READING: "Currently Reading",
  FINISHED: "Finished"
};

export default function ShelfPage() {
  const { token } = useAuth();
  const [records, setRecords] = useState<ShelfRecord[]>([]);
  const [books, setBooks] = useState<BookSummary[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const fetchShelf = async () => {
    setLoading(true);
    setError(null);
    try {
      const [shelfResponse, booksResponse] = await Promise.all([
        fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/shelf`, {
          headers: {
            Authorization: token ? `Bearer ${token}` : ""
          }
        }),
        fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/books`, {
          headers: {
            Authorization: token ? `Bearer ${token}` : ""
          }
        })
      ]);

      if (!shelfResponse.ok) {
        setError("Unable to load shelf.");
        return;
      }

      if (!booksResponse.ok) {
        setError("Unable to load books.");
        return;
      }

      const shelfData: ShelfRecord[] = await shelfResponse.json();
      const bookData: BookSummary[] = await booksResponse.json();
      setRecords(shelfData);
      setBooks(bookData);
    } catch (err) {
      setError("Unable to load shelf.");
    } finally {
      setLoading(false);
    }
  };

  const getBookSummary = (bookId: number) => {
    return books.find((book) => book.id === bookId);
  };

  const handleDelete = async (recordId: number) => {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/shelf/${recordId}`, {
        headers: {
          Authorization: token ? `Bearer ${token}` : ""
        }
      });
      if (!response.ok) {
        setError("Unable to remove book.");
        return;
      }

      setRecords((prev) => prev.filter((record) => record.id !== recordId));
    } catch (err) {
      setError("Unable to remove book.");
    }
  };

  useEffect(() => {
    fetchShelf();
  }, []);

  const grouped = useMemo(() => {
    return records.reduce(
      (acc, record) => {
        acc[record.status].push(record);
        return acc;
      },
      {
        WANT_TO_READ: [] as ShelfRecord[],
        READING: [] as ShelfRecord[],
        FINISHED: [] as ShelfRecord[]
      }
    );
  }, [records]);

  const renderStars = (rating?: number | null) => {
    if (!rating) {
      return "";
    }
    return "★★★★★".slice(0, Math.min(rating, 5));
  };

  return (
    <main className="mx-auto max-w-4xl p-6">
      <h1 className="text-2xl font-semibold">My Shelf</h1>
      <p className="mt-2 text-slate-600">
        Track what you want to read, what you're currently reading, and what you've finished.
      </p>

      {error ? <p className="mt-4 text-sm text-red-600">{error}</p> : null}
      {loading ? <p className="mt-4 text-sm text-slate-600">Loading shelf...</p> : null}

      <div className="mt-6 space-y-6">
        {(Object.keys(statusLabels) as ShelfRecord["status"][]).map((status) => (
          <section key={status} className="rounded border border-slate-200 bg-white p-4">
            <h2 className="text-lg font-semibold">{statusLabels[status]}</h2>
            {grouped[status].length === 0 ? (
              <p className="mt-2 text-sm text-slate-500">No books yet.</p>
            ) : (
              <ul className="mt-3 space-y-3">
                {grouped[status].map((record) => (
                  <li key={record.id} className="flex flex-col gap-2 border-b border-slate-100 pb-3 last:border-b-0">
                    <div className="flex items-start justify-between">
                      <div>
                        <p className="font-medium">
                          {getBookSummary(record.bookId)?.title ?? "Book"}
                        </p>
                        <p className="text-sm text-slate-600">
                          {getBookSummary(record.bookId)?.author ?? ""}
                        </p>
                        <p className="text-xs uppercase text-slate-400">{statusLabels[record.status]}</p>
                      </div>
                      <button
                        className="text-sm text-red-600"
                        type="button"
                        onClick={() => handleDelete(record.id)}
                      >
                        Remove
                      </button>
                    </div>
                    {record.status === "FINISHED" ? (
                      <div className="text-sm text-slate-600">
                        {record.rating ? <p>Rating: {renderStars(record.rating)}</p> : null}
                        {record.review ? <p className="mt-1">{record.review}</p> : null}
                      </div>
                    ) : null}
                  </li>
                ))}
              </ul>
            )}
          </section>
        ))}
      </div>
    </main>
  );
}
