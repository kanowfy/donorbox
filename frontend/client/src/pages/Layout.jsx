import { Outlet, Link } from "react-router-dom";
import Header from "../components/Header";
import Footer from "../components/Footer";
import { Banner } from "flowbite-react";
import { MdAnnouncement } from "react-icons/md";
import { HiX } from "react-icons/hi";
import { useAuthContext } from "../context/AuthContext";

const Layout = () => {
  const { user } = useAuthContext();
  return (
    <div>
      {user?.verification_status === "unverified" && (
        <Banner>
          <div className="flex w-full justify-between border-b border-gray-200 bg-gray-50 p-4 dark:border-gray-600 dark:bg-gray-700">
            <div className="mx-auto flex items-center">
              <p className="flex items-center text-sm font-normal text-gray-500 dark:text-gray-400">
                <MdAnnouncement className="mr-4 h-4 w-4" />
                <span className="[&_p]:inline">
                  You have not verified your account. Please follow this link to
                  verify your account so you can start your own campaigns.{" "}
                  <Link
                    to="/account/verify"
                    className="inline font-medium text-cyan-600 underline decoration-solid underline-offset-2 hover:no-underline dark:text-cyan-500"
                  >
                    Verify
                  </Link>
                </span>
              </p>
            </div>
            <Banner.CollapseButton
              color="gray"
              className="border-0 bg-transparent text-gray-500 dark:text-gray-400"
            >
              <HiX className="h-4 w-4" />
            </Banner.CollapseButton>
          </div>
        </Banner>
      )}
      <div>
        <Header />
      </div>
      <div className="flex flex-col min-h-screen">
        <Outlet />
      </div>
      <div className="px-6 mt-auto">
        <Footer />
      </div>
    </div>
  );
};

export default Layout;
