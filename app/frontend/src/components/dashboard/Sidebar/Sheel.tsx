import { SidebarProvider } from "./SidebarContext";
import Sidebar from "./index";
export const Shell = () => {
    return (
        <SidebarProvider>
            <Sidebar />
        </SidebarProvider>
    );
};