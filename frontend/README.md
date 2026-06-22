# Vayana Frontend

Simple Next.js (App Router) frontend for the Vayana book reading tracker.

## Routes
- `/login`
- `/register`
- `/` (protected)
- `/shelf` (protected)
- `/books` (protected)

## Auth
A lightweight `AuthContext` stores the JWT token in `localStorage` and a cookie named `vayana_token` so middleware can protect routes.
