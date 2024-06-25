"use client";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

const ProtectedRoute = ({ children }) => {
  const router = useRouter();
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  useEffect(() => {
    if (isClient) {
      const user = localStorage.getItem("user");
      if (!user) {
        router.push("/login");
      }
    }
  }, [isClient, router]);

  if (!isClient) {
    return null; // or a loading spinner
  }

  return children;
};

export default ProtectedRoute;
