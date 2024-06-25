
import SideNavBar from "@/components/Sidebar/SideNavBar";
import Main from "@/components/Main/Main";
import ProtectedRoute from "@/components/ProtectedRoutes/ProctectedRoutes";
import Storage from "@/components/Storage/Storage";

export default function Home() {
  
  return (
    <ProtectedRoute>
        <div className="flex">
          <SideNavBar />
          <div className="grid grid-cols-1 md:grid-cols-3 w-full">
            <div className="col-span-2 p-5">
              <Main />
            </div>
            <div className="bg-white p-5 order-first md:order-last">
              <Storage/>
            </div>
          </div>
        </div>
    </ProtectedRoute>
  );
}
