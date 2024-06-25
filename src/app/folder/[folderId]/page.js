"use client";
import React, { useContext, useState, useEffect } from "react";
import { useRouter ,useSearchParams } from "next/navigation";
import { deleteDoc, doc, getFirestore, collection, query, where, getDocs } from "firebase/firestore";
import { app } from "@/firebase/firebase";
import SideNavBar from "@/components/Sidebar/SideNavBar";
import FileList from "@/components/Main/File/FileList";
import { ShowToastContext } from "@/context/ShowToastContext";
import Main from "@/components/Main/Main";
import Storage from "@/components/Storage/Storage";

function Page() {
  const searchParams = useSearchParams();
  const name = searchParams.get('name');
  const id = searchParams.get('id');
  const [fileList, setFileList] = useState([]);
  const { setShowToastMsg } = useContext(ShowToastContext);
  const router = useRouter();
  const db = getFirestore(app);

  const deleteFolder = async () => {
    try {
      await deleteDoc(doc(db, "Folders", id));
      setShowToastMsg('Folder Deleted!');
      router.back();
    } catch (err) {
      console.error(err);
      setShowToastMsg('Error deleting folder');
    }
  };

  return (
    <div className="flex">
      <SideNavBar />
      <div className="grid grid-cols-1 md:grid-cols-3 w-full">
        <div className="col-span-2 p-5">
          <h2 className='text-[20px] font-bold mt-5'>{name}
            <svg xmlns="http://www.w3.org/2000/svg"
              onClick={() => deleteFolder()}
              fill="none" viewBox="0 0 24 24"
              strokeWidth={1.5} stroke="currentColor"
              className="w-5 h-5 float-right text-red-500 hover:scale-110 transition-all cursor-pointer">
              <path strokeLinecap="round" strokeLinejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
            </svg>
          </h2>
          <Main/>
        </div>
        <div className="bg-white p-5 order-first md:order-last">
          <Storage/>
        </div>
      </div>
    </div>
  );
}

export default Page;
