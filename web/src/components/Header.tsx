"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { logoutUser } from "@/api/auth";
import { User } from "@/types/user";
import { Avatar } from "./Avatar";
import { Button } from "./Button";

export type HeaderProps = {
  withAuth: boolean;
  user?: User | null;
};

export const Header: React.FC<HeaderProps> = ({ withAuth, user }) => {
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
      {withAuth && user && (
        <div className="flex items-center gap-2">
          <span>{user.login}</span>
          {user.avatar && <Avatar src={user.avatar} alt="" size={42} />}
          <Button className="ml-4" onClick={handleLogout}>
            Logout
          </Button>
        </div>
      )}
    </header>
  );
};
