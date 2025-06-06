-- CreateEnum
CREATE TYPE "Status" AS ENUM ('active', 'deactived', 'pending');

-- CreateTable
CREATE TABLE "users" (
    "id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "status" "Status" NOT NULL DEFAULT 'pending',

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);
