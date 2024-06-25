"use client";
import { useState, useEffect } from "react";
import { signOut } from "firebase/auth";
import Image from "next/image";
import { auth } from "@/firebase/firebase"; // Ensure auth is imported correctly

function UserInfo() {
  const [user, setUser] = useState(null);
  const [showToastMsg, setShowToastMsg] = useState("");

  useEffect(() => {
    const storedUser = JSON.parse(localStorage.getItem("user"));
    if (storedUser) {
      setUser(storedUser);
    }
  }, []);

  const logout = async () => {
    try {
      await signOut(auth);
      localStorage.removeItem("user");
      setShowToastMsg("Logout successful!");
    } catch (err) {
      console.error(err);
      setShowToastMsg("Error logging out: " + err.message);
    }
  };

  return (
    <div>
      {user ? (
        <div className='flex gap-2 items-center'>
          <Image
            src={user.photoURL}
            alt='user-image'
            width={40}
            height={40}
            className='rounded-full'
          />
          <div>
            <h2 className='text-[15px] font-bold'>{user.displayName}</h2>
            <h2 className='text-[13px] text-gray-400 mt-[-4px]'>{user.email}</h2>
          </div>
          <div className='bg-blue-200 p-2 rounded-lg cursor-pointer' onClick={logout}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-6 h-6 text-blue-500 hover:animate-pulse transition-all"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15M12 9l-3 3m0 0l3 3m-3-3h12.75"
              />
            </svg>
          </div>
        </div>
      ) : null}
      {showToastMsg && <div>{showToastMsg}</div>}
    </div>
  );
}

export default UserInfo;
