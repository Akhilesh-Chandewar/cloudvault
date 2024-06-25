import React from "react";

function StorageDetailItem({ item = {} }) {
  const { type = "", totalFile = "", size = "", logo = "" } = item;

  return (
    <>
      <div className='flex justify-between items-center mt-3 border-b-[1px] pb-2'>
        <div className='flex items-center gap-4'>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            className={`w-8 h-8 
            rounded-md
            p-2 ${
              type === 'Videos' ? 'bg-blue-200 text-blue-600' :
              type === 'Documents' ? 'bg-yellow-200 text-yellow-600' :
              type === 'Others' ? 'bg-red-200 text-red-600' : 'bg-green-200 text-green-600'
            } `}
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d={logo}
            />
          </svg>
          <div>
            <h2 className='text-[14px] font-semibold'>{type}</h2>
            <h2 className='text-[12px] mt-[-4px] text-gray-400'>{totalFile} Files</h2>
          </div>
        </div>
        <div className='font-semibold'>{size}</div>
      </div>
    </>
  );
}

export default StorageDetailItem;
