"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { logoutUser } from "@/api/auth";
import { Button } from "./Button";

export type HeaderProps = {
  withAuth: boolean;
};

export const Header: React.FC<HeaderProps> = ({ withAuth }) => {
  const router = useRouter();

  const handleLogout = () => {
    logoutUser()
      .then(() => {
        router.push("/login");
        router.refresh();
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <header className="w-full flex justify-between items-center gap-2">
      <h1 className="text-2xl">cronjobs</h1>
      {withAuth && <Button onClick={handleLogout}>Logout</Button>}
    </header>
  );
};
