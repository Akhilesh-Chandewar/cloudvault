import { Inter } from "next/font/google";
import "./globals.css";
import { ShowToastProvider } from "@/context/ShowToastContext";
import Toast from "@/components/Notification/Toast";
import {  ParentFolderIdProvider } from '@/context/ParentFolderIdContext';

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Cloud Vault",
  description: "A cloud based file management software",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ShowToastProvider>
          <ParentFolderIdProvider>
            {children}
            <Toast />
          </ParentFolderIdProvider>
        </ShowToastProvider>
      </body>
    </html>
  );
}
