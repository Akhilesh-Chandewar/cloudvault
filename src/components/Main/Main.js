"use client";
import { useState, useEffect, useContext } from "react";
import { useSearchParams } from "next/navigation";
import FolderList from "./Folder/FolderList";
import FileList from "./File/FileList";
import SearchBar from "./SearchBar";
import { getFirestore, collection, query, where, getDocs } from "firebase/firestore";
import { app } from "@/firebase/firebase";
import { ParentFolderIdContext } from "@/context/ParentFolderIdContext";

function Main() {
  const searchParams = useSearchParams();
  const parentFolderId = searchParams.get('id') || 0; // Default to "0" if id is not provided
  const [folderList, setFolderList] = useState([]);
  const [fileList, setFileList] = useState([]);
  const [loading, setLoading] = useState(true);
  const db = getFirestore(app);
  const { setParentFolderId } = useContext(ParentFolderIdContext);

  const getFolderList = async (parentId) => {
    setFolderList([]);
    const user = JSON.parse(localStorage.getItem("user"));
    if (!user) return;

    const q = query(
      collection(db, "Folders"),
      where("createdBy", "==", user.email)
    );

    const querySnapshot = await getDocs(q);
    const folders = [];
    querySnapshot.forEach((doc) => {
      folders.push(doc.data());
    });
    const f = folders.filter((folder)=>{
      folder.parentFolderId === parentId
    })
    console.log(f)
    setFolderList(folders);
  };

  const getFileList = async (parentId) => {
    setFileList([]);
    const user = JSON.parse(localStorage.getItem("user"));
    if (!user) return;

    const q = query(
      collection(db, "Files"),
      where("createdBy", "==", user.email)
    );

    const querySnapshot = await getDocs(q);
    const files = [];
    querySnapshot.forEach((doc) => {
      files.push(doc.data());
    });
    setFileList(files);
  };

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      try {
        await getFolderList(parentFolderId);
        await getFileList(parentFolderId);
      } catch (error) {
        console.error("Error fetching data: ", error);
      } finally {
        setLoading(false);
      }
    };

    setParentFolderId(Number(parentFolderId));
    fetchData();
  }, [parentFolderId]);

  console.log(fileList)
  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <SearchBar />
      <FolderList folderList={folderList} />
      <FileList fileList={fileList} />
    </div>
  );
}

export default Main;
