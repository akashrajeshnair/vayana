"use client";

import type { ChangeEvent, FormEvent } from "react";
import { useState } from "react";
import { useAuth } from "@/context/AuthContext";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  const { login } = useAuth();
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setError(null);
    if (!email || !password) {
      return;
    }
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/auth/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ email, password })
      });

      if (!response.ok) {
        setError("Login failed. Please check your credentials.");
        return;
      }

      const data: { token?: string } = await response.json();
      if (!data.token) {
        setError("Login failed. Missing token in response.");
        return;
      }

      login(data.token);
      router.push("/");
    } catch (err) {
      setError("Login failed. Please try again.");
    }
  };

  return (
    <main className="mx-auto max-w-md p-6">
      <h1 className="text-2xl font-semibold">Login</h1>
      <form className="mt-4 space-y-4" onSubmit={handleSubmit}>
        <label className="block">
          <span className="text-sm text-slate-600">Email</span>
          <input
            className="mt-1 w-full rounded border border-slate-300 p-2"
            type="email"
            value={email}
            onChange={(event: ChangeEvent<HTMLInputElement>) => setEmail(event.target.value)}
            required
          />
        </label>
        <label className="block">
          <span className="text-sm text-slate-600">Password</span>
          <input
            className="mt-1 w-full rounded border border-slate-300 p-2"
            type="password"
            value={password}
            onChange={(event: ChangeEvent<HTMLInputElement>) => setPassword(event.target.value)}
            required
          />
        </label>
        <button
          className="w-full rounded bg-slate-900 px-4 py-2 text-white"
          type="submit"
        >
          Sign in
        </button>
        {error ? (
          <p className="text-sm text-red-600">{error}</p>
        ) : null}
      </form>
    </main>
  );
}
