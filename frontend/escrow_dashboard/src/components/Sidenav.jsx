import { Sidebar, Avatar } from "flowbite-react";
import { Link } from "react-router-dom";
import { useAuthContext } from "../context/AuthContext";
import { LuLogOut } from "react-icons/lu";
import { RxDashboard } from "react-icons/rx";
import { MdPayments } from "react-icons/md";
import { RiRefund2Fill } from "react-icons/ri";
import { AiOutlineAudit } from "react-icons/ai";

const customSidenavTheme = {
  root: {
    base: "h-full",
    collapsed: {
      on: "w-16",
      off: "w-64",
    },
    inner: "h-full overflow-y-auto overflow-x-hidden bg-gray-800 px-3 py-4",
  },
  collapse: {
    button:
      "group flex w-full items-center rounded-lg p-2 text-base font-normal transition duration-75 text-white hover:bg-gray-700",
    icon: {
      base: "h-6 w-6 transition duration-75 text-gray-400 group-hover:text-white",
      open: {
        off: "",
        on: "text-gray-900",
      },
    },
    label: {
      base: "ml-3 flex-1 whitespace-nowrap text-left",
      icon: {
        base: "h-6 w-6 transition delay-0 ease-in-out",
        open: {
          on: "rotate-180",
          off: "",
        },
      },
    },
    list: "space-y-2 py-2",
  },
  cta: {
    base: "mt-6 rounded-lg p-4 bg-gray-700",
    color: {
      blue: "bg-cyan-900",
      dark: "bg-dark-900",
      failure: "bg-red-900",
      gray: "bg-alternative-900",
      green: "bg-green-900",
      light: "bg-light-900",
      red: "bg-red-900",
      purple: "bg-purple-900",
      success: "bg-green-900",
      yellow: "bg-yellow-900",
      warning: "bg-yellow-900",
    },
  },
  item: {
    base: "flex items-center justify-center rounded-lg p-2 text-base font-normal text-white hover:bg-gray-700",
    active: "bg-gray-700",
    collapsed: {
      insideCollapse: "group w-full pl-8 transition duration-75",
      noIcon: "font-bold",
    },
    content: {
      base: "flex-1 whitespace-nowrap px-3",
    },
    icon: {
      base: "h-6 w-6 flex-shrink-0 transition duration-75 text-gray-400 group-hover:text-white",
      active: "text-gray-100",
    },
    label: "",
    listItem: "",
  },
  items: {
    base: "",
  },
  itemGroup: {
    base: "mt-4 space-y-2 border-t pt-4 first:mt-0 first:border-t-0 first:pt-0 border-gray-700",
  },
  logo: {
    base: "mb-5 flex items-center pl-2.5",
    collapsed: {
      on: "hidden",
      off: "self-center whitespace-nowrap text-xl font-semibold text-white",
    },
    img: "mr-3 h-6 sm:h-7",
  },
};

const Sidenav = () => {
  const { user } = useAuthContext();
  return (
    <Sidebar
      className="h-screen sticky top-0 w-full"
      theme={customSidenavTheme}
    >
      <div className="flex flex-col justify-between h-full">
        <div>
          <div className="flex justify-center pr-5 pt-3">
            <Sidebar.Logo
              as={Link}
              to="#"
              img="/logo.svg"
              imgAlt="Donorbox logo"
            >
              Donorbox
            </Sidebar.Logo>
          </div>
          <div className="mt-10">
            <div className="text-gray-500 text-sm ml-2 mb-2">MENU</div>
            <Sidebar.Items>
              <Sidebar.ItemGroup>
                <Sidebar.Item as={Link} to="/" icon={RxDashboard}>
                  Dashboard
                </Sidebar.Item>
                <Sidebar.Item as={Link} to="/manage/payout" icon={MdPayments}>
                  Manage Payouts
                </Sidebar.Item>
                <Sidebar.Item
                  as={Link}
                  to="/manage/refund"
                  icon={RiRefund2Fill}
                >
                  Manage Refunds
                </Sidebar.Item>
                <Sidebar.Item
                  as={Link}
                  to="/transactions"
                  icon={AiOutlineAudit}
                >
                  Transaction Audits
                </Sidebar.Item>
              </Sidebar.ItemGroup>
              <Sidebar.ItemGroup>
                <Sidebar.Item as={Link} to="/logout" icon={LuLogOut}>
                  Logout
                </Sidebar.Item>
              </Sidebar.ItemGroup>
            </Sidebar.Items>
          </div>
        </div>
        <div className="w-full flex justify-center space-x-2">
          <Avatar
            img={() => (
              <img
                src="/avatar.svg"
                className="h-8 w-8 bg-white rounded-full p-1"
              />
            )}
            rounded
            bordered
          />
          <div className="text-gray-300 text-sm mt-1">{user?.email}</div>
        </div>
      </div>
    </Sidebar>
  );
};

export default Sidenav;
