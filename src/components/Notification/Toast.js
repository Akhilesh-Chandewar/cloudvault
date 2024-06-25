"use client";
import React, { useContext, useEffect } from "react";
import { ShowToastContext } from "@/context/ShowToastContext";

function Toast() {
  const { showToastMsg, setShowToastMsg } = useContext(ShowToastContext);

  useEffect(() => {
    if (showToastMsg) {
      const timer = setTimeout(() => {
        setShowToastMsg(null);
      }, 3000);
      return () => clearTimeout(timer); // Clear the timeout if component unmounts
    }
  }, [showToastMsg, setShowToastMsg]);

  if (!showToastMsg) return null;

  return (
    <div className="toast toast-top toast-end">
      <div className="alert alert-success">
        <span>{showToastMsg}</span>
      </div>
    </div>
  );
}

export default Toast;
