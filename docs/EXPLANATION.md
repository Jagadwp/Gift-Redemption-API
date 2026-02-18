# Database And Architecture Notes

This document accompanies the database dump and ERD image in `docs/`.

## Physical ERD (Schema) Explanation

ERD image: `docs/gift_redemption - public.png`

![ERD](<gift_redemption - public.png>)

**Entities and relationships**

- `users` (1) --- (N) `redemptions`
- `gifts` (1) --- (N) `redemptions`
- `users` (1) --- (N) `ratings`
- `gifts` (1) --- (N) `ratings`
- `redemptions` (1) --- (1) `ratings` via `ratings.redemption_id` (unique)

**Why this structure**

- `redemptions` is the transaction log between a user and a gift. It preserves quantity and total points at the time of redeem.
- `ratings` is tied to a specific redemption to enforce "one rating per redemption" (unique constraint on `redemption_id`).
- `gifts` stores aggregate fields (`avg_rating`, `total_reviews`) for fast list and sorting queries.
- `users` has soft delete support (`deleted_at`) to preserve history while hiding inactive users.

**Key columns**

- `gifts.avg_rating` (numeric(3,2)) stores aggregate rating.
- `ratings.score` (numeric(2,1)) allows half-star ratings (1.0 to 5.0).
- `redemptions.total_point` captures points paid at redemption time.

## Clean Architecture

This project uses a layered Clean Architecture:

- **Handler (HTTP layer)**: request parsing, validation, and response formatting.
- **Service (use-case layer)**: business rules, transactions, and orchestration.
- **Repository (data access layer)**: isolated database operations (CRUD + queries).
- **Model (domain entities)**: GORM models that map to DB tables.

This separation keeps business logic independent from transport and data storage, making the codebase easier to test and evolve.

## Why PostgreSQL (vs MySQL)

PostgreSQL is chosen because:

- Strong correctness guarantees for transactions and constraints (important for stock deduction + redemption workflow).
- Robust numeric types and check constraints for rating validation.
- Better support for advanced indexing and partial indexes used in this schema.
- Mature tooling for backups and schema dumps (`pg_dump`) which is used in this repo.

MySQL would also work, but PostgreSQL provides stricter data integrity and richer indexing options for this use case.

## Database Optimization Applied

Based on the schema dump (`docs/dump_gift_redemption_schema.sql`), these optimizations are present:

- Index: `idx_users_email` (partial index on `email` for non-deleted users)
- Index: `idx_gifts_created_at` and `idx_gifts_avg_rating` for listing/sorting
- Index: `idx_redemptions_user_id`, `idx_redemptions_gift_id` for lookup by user/gift
- Index: `idx_ratings_user_id`, `idx_ratings_gift_id` for aggregate stats
- Soft delete support via `deleted_at` and indexes on that column to keep queries fast.
- Aggregated fields in `gifts` (`avg_rating`, `total_reviews`) to avoid heavy recalculation on every list query.
- Constraint: `ratings.score` check constraint (1 to 5)
- Constraint: `uq_rating_redemption` unique constraint ensures one rating per redemption

These choices reduce query cost for common endpoints (`/gifts` listing, `/gifts/:id`, and user activity history) while preserving data integrity.

## Deployment (Heroku)

Heroku was chosen for **simplicity and speed**: it handles build, run, and PostgreSQL provisioning with minimal setup, which keeps the review process lightweight. For a backend API like this, Heroku offers a practical path to production without extra infrastructure complexity.

## Files in `docs/`

- `docs/dump_gift_redemption_schema.rar`: compressed database export (contains `.sql` schema + data)
- `docs/gift_redemption - public.png`: physical ERD image
- `docs/Gift Redemption API.postman_collection.json`: Postman collection

Published Postman doc URL:
https://documenter.getpostman.com/view/12534314/2sBXcDGM9Z
