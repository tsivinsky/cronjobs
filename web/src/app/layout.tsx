import { cookies } from "next/headers";
import { Header } from "@/components/Header";
import { LayoutProps } from "@/types/next";
import "./globals.css";

export const metadata = {
  title: "cronjobs",
};

export default function RootLayout({ children }: LayoutProps) {
  const c = cookies();
  const userId = c.get("user_id")?.value;

  const hasAuth = userId !== undefined;

  return (
    <html lang="en">
      <body className="max-w-screen-2xl mx-auto p-2 min-h-screen bg-bg">
        <Header withAuth={hasAuth} />
        <main>{children}</main>
      </body>
    </html>
  );
}
