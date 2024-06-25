"use client"
import { useState } from "react";
import FolderItem from "./FolderItem";
import { useRouter } from "next/navigation";
import FolderItemSmall from "./FolderItemSmall";

function FolderList({ folderList, isBig = true }) {
  const [activeFolder, setActiveFolder] = useState();
  const router = useRouter();

  const onFolderClick = (index, item) => {
    setActiveFolder(index);
    router.push(`/folder/${item.id}?name=${item.name}&id=${item.id}`);
  };

  return (
    <div className="p-5 mt-5 bg-white rounded-lg">
      {isBig && (
        <h2 className="text-17px font-bold items-center">
          Recent Folders
        </h2>
      )}
      <div className={isBig ? "grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 mt-3 gap-4" : ""}>
        {folderList.map((item, index) => (
          <div key={index} onClick={() => onFolderClick(index, item)}>
            {isBig ? <FolderItem folder={item} /> : <FolderItemSmall folder={item} />}
          </div>
        ))}
      </div>
    </div>
  );
}

export default FolderList;
