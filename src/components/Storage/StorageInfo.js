"use client";
import { collection, getDocs, getFirestore, query, where } from 'firebase/firestore';
import React, { useEffect, useState } from 'react';
import { app } from '@/firebase/firebase';
import StorageSize from '@/service/StorageSize';

function StorageInfo() {
  const db = getFirestore(app);
  const [totalSizeUsed, setTotalSizeUsed] = useState(0);
  const [imageSize, setImageSize] = useState(0);
  const user = JSON.parse(localStorage.getItem("user"));
  const [fileList, setFileList] = useState([]);
  let totalSize = 0;

  useEffect(() => {
    if (user) {
      totalSize = 0;
      getAllFiles();
    }
  }, [user]);

  useEffect(() => {
    if (fileList.length > 0) {
      setImageSize(StorageSize.getStorageByType(fileList, ['png', 'jpg']));
    }
  }, [fileList]);

  const getAllFiles = async () => {
    const q = query(
      collection(db, "files"),
      where("createdBy", "==", user.email)
    );
    const querySnapshot = await getDocs(q);
    const files = [];
    querySnapshot.forEach((doc) => {
      totalSize += doc.data().size;
      files.push(doc.data());
    });
    setFileList(files);
    setTotalSizeUsed((totalSize / 1024 ** 2).toFixed(2) + " MB");
  };

  return (
    <div className='mt-7'>
      <h2 className="text-[22px] font-bold">
        {totalSizeUsed} {" "}
        <span className="text-[14px] font-medium">
          used of{" "}
        </span>{" "}
        50 MB
      </h2>
      <div className='w-full bg-gray-200 h-2.5 flex'>
        <div className='bg-blue-600 h-2.5 w-[25%]'></div>
        <div className='bg-green-600 h-2.5 w-[35%]'></div>
        <div className='bg-yellow-400 h-2.5 w-[15%]'></div>
      </div>
    </div>
  );
}

export default StorageInfo;
