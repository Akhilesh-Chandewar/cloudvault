"use client";
import Image from "next/image";
import React, { useState, useContext } from "react";
import { app } from "@/firebase/firebase";
import { doc, getFirestore, setDoc } from "firebase/firestore";
import { ShowToastContext } from "@/context/ShowToastContext";
import { ParentFolderIdContext , setParentFolderId } from "@/context/ParentFolderIdContext";

function CreateFolderModal() {
  const docId = Date.now().toString();
  const [folderName, setFolderName] = useState("");
  const db = getFirestore(app);
  const { setShowToastMsg } = useContext(ShowToastContext);
  const {parentFolderId,setParentFolderId}=useContext(ParentFolderIdContext)
  const onCreate = async () => {
    const user = JSON.parse(localStorage.getItem("user"));
    if (!folderName) {
      alert("Folder name is required");
      return;
    }
    try {
      await setDoc(doc(db, "Folders", docId), {
        name: folderName,
        id: docId,
        createdBy: user.email,
        parentFolderId:parentFolderId
      });
      window.location.reload();
      setShowToastMsg("Folder Created!");
    } catch (err) {
      console.error(err);
      setShowToastMsg("Error creating folder");
    }
  };

  return (
    <div>
      <form method="dialog" className="modal-box p-9 items-center">
        <button className="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">
          ✕
        </button>
        <div className="w-full items-center flex flex-col justify-center gap-3">
          <Image src="/folder.png" alt="folder" width={50} height={50} />
          <input
            type="text"
            placeholder="Folder Name"
            className="p-2 border-[1px] outline-none rounded-md"
            onChange={(e) => setFolderName(e.target.value)}
          />
          <button
            type="button"
            className="bg-blue-500 text-white rounded-md p-2 px-3 w-full"
            onClick={onCreate}
          >
            Create
          </button>
        </div>
      </form>
    </div>
  );
}

export default CreateFolderModal;
