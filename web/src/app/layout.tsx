import { Header } from "@/components/Header";
import { LayoutProps } from "@/types/next";
import "./globals.css";
import { getServerUser } from "@/api/user";

export const metadata = {
  title: "cronjobs",
};

export default async function RootLayout({ children }: LayoutProps) {
  const user = await getServerUser();
  const hasAuth = !!user;

  return (
    <html lang="en">
      <body className="max-w-screen-2xl mx-auto p-2 min-h-screen bg-bg">
        <Header withAuth={hasAuth} user={user} />
        <main>{children}</main>
      </body>
    </html>
  );
}
