"use client"
import { createContext, useState } from "react";

export const ShowToastContext = createContext();

export const ShowToastProvider = ({ children }) => {
  const [showToastMsg, setShowToastMsg] = useState(null);

  return (
    <ShowToastContext.Provider value={{ showToastMsg, setShowToastMsg }}>
      {children}
    </ShowToastContext.Provider>
  );
};
