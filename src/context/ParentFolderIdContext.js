"use client"
import React, { createContext, useState } from 'react';

export const ParentFolderIdContext = createContext();

export const ParentFolderIdProvider = ({ children }) => {
  const [parentFolderId, setParentFolderId] = useState(0);

  return (
    <ParentFolderIdContext.Provider value={{ parentFolderId, setParentFolderId }}>
      {children}
    </ParentFolderIdContext.Provider>
  );
};
