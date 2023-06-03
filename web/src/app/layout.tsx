import { LayoutProps } from "@/types/next";
import "./globals.css";

export const metadata = {
  title: "cronjobs",
};

export default function RootLayout({ children }: LayoutProps) {
  return (
    <html lang="en">
      <body className="max-w-screen-2xl mx-auto p-2 min-h-screen bg-bg">
        <main>{children}</main>
      </body>
    </html>
  );
}
