"use client";
import { auth, googleProvider } from "@/firebase/firebase";
import {
  createUserWithEmailAndPassword,
  signInWithPopup,
  signOut,
} from "firebase/auth";
import { useState, useEffect, useContext } from "react";
import { useRouter } from "next/navigation";
import { ShowToastContext } from "@/context/ShowToastContext";

function Auth() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { setShowToastMsg } = useContext(ShowToastContext);
  const router = useRouter();

  useEffect(() => {
    const user = localStorage.getItem("user");
    if (user) {
      router.push("/");
    }
  }, [router]);

  const signUp = async () => {
    try {
      const userCredential = await createUserWithEmailAndPassword(auth, email, password);
      localStorage.setItem("user", JSON.stringify(userCredential.user));
      setShowToastMsg("Sign up successful!");
      router.push("/");
    } catch (err) {
      console.error(err);
      setShowToastMsg("Error signing up: " + err.message);
    }
  };

  const signInWithGoogle = async () => {
    try {
      const result = await signInWithPopup(auth, googleProvider);
      localStorage.setItem("user", JSON.stringify(result.user));
      setShowToastMsg("Sign in with Google successful!");
      router.push("/");
    } catch (err) {
      console.error(err);
      setShowToastMsg("Error signing in with Google: " + err.message);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-xs p-6 bg-white rounded-md shadow-md">
        <h1 className="mb-4 text-xl font-semibold text-center">Auth</h1>
        <input
          className="w-full px-3 py-2 mb-4 text-sm border rounded shadow-sm border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Email..."
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          className="w-full px-3 py-2 mb-4 text-sm border rounded shadow-sm border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Password..."
          type="password"
          onChange={(e) => setPassword(e.target.value)}
        />
        <button
          className="w-full px-4 py-2 mb-2 text-white bg-blue-500 rounded hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
          onClick={signUp}
        >
          Sign Up
        </button>
        <button
          className="w-full px-4 py-2 mb-2 text-white bg-red-500 rounded hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500"
          onClick={signInWithGoogle}
        >
          Sign In With Google
        </button>
      </div>
    </div>
  );
}

export default Auth;
