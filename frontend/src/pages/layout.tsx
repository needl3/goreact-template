import { CheckmarkIcon, Toaster } from "react-hot-toast";
import { Outlet } from "react-router-dom";
import IconInfoCircle from "../components/icons/info-circle";
import { IconLoadingCircle } from "../components/icons/loading-circle";
import IconCross from "../components/icons/cross";
import { routes } from "../routes";
import { NavLink } from "react-router-dom";
import { CustomButton } from "../components/button";

export function Layout() {
  const renderTopNav = <div className="w-full flex items-center py-2 px-5 bg-gray-50 justify-between">
    <div className="flex gap-5 items-center justify-start">
      <h1 className="text-lg font-bold flex gap-3 bg-primary-200 p-3 uppercase rounded-bl-xl rounded-tr-xl">GoReact App</h1>
      <span>A ready to use template with Golang ❤️ React</span>
    </div>
    <CustomButton name="Logout" link={routes.api.auth.signout} />
  </div>

  const renderNavItem = (link: string, text: string) => <NavLink to={link} className={({ isActive }) => `${isActive ? "bg-primary-100" : ""} py-2 px-5 duration-500 rounded-tr-xl rounded-bl-xl`}>{text}</NavLink>

  const renderNav = <div className="w-[300px] h-full flex flex-col gap-3 border-r p-5">
    {renderNavItem(routes.pages.new, "New")}
  </div>

  const renderToaster = <Toaster
    position="top-right"
    reverseOrder={false}
    gutter={8}
    toastOptions={{
      duration: 5000,
      icon: <IconInfoCircle fill="#0A175B" />,
      style: {
        background: '#E7F2F1',
        color: '#0A175B',
      },
      loading: {
        style: {
          color: '#0A175B',
          backgroundColor: '#E7F2F1',
        },
        icon: (
          <div className="animate-spin">
            <IconLoadingCircle stroke="#0A175B" className="w-4 h-4" />
          </div>
        ),
        duration: 2000,
      },
      success: {
        style: { backgroundColor: '#BAFF26', color: 'black' },
        duration: 2000,
        icon: <CheckmarkIcon />,
      },
      error: {
        style: { backgroundColor: 'red', color: 'white' },
        duration: 2000,
        icon: <IconCross fill="white" />,
      },
    }}
  />

  return <div className="h-dvh w-full border-red-500 flex flex-col">
    {renderTopNav}
    <div className="flex w-full h-full">
      {renderNav}
      <div className="flex items-center justify-center w-full h-full p-10">
        <Outlet />
        {renderToaster}
      </div>
    </div>
  </div>
}
