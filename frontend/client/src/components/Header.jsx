import { MegaMenu, Dropdown, Navbar, Avatar, Button } from "flowbite-react";
import { useAuthContext } from "../context/AuthContext";
import { Link } from "react-router-dom";
import Notification from "./Notification";
import { BiPlusMedical } from "react-icons/bi";
import { FaTruckMedical } from "react-icons/fa6";
import { FaGraduationCap } from "react-icons/fa";
import { PiDogFill } from "react-icons/pi";
import { MdEmojiEvents } from "react-icons/md";
import { MdEventAvailable } from "react-icons/md";
import { FaTree } from "react-icons/fa";
import { GiSailboat } from "react-icons/gi";
import { MdBusinessCenter } from "react-icons/md";
import { useEffect, useState } from "react";
import projectService from "../services/project";

const Header = () => {
  const { user } = useAuthContext();

  return (
    <MegaMenu
      className="flex justify-center w-full"
      theme={{
        root: {
          base: "bg-green-50 dark:border-gray-700 dark:bg-gray-800 sm:px-4",
          inner: {
            base: "items-center justify-between w-full",
          },
        },
      }}
    >
      <div className="w-full flex items-center justify-between py-4 px-10 md:space-x-8 border-b border-gray-300">
        <Navbar.Brand as={Link} to="/" className="flex-grow basis-0">
          <img
            src="/logo.svg"
            className="mr-3 h-6 sm:h-9"
            alt="Flowbite React Logo"
          />
          <span className="self-center whitespace-nowrap text-2xl font-semibold tracking-tighter">
            Donorbox
          </span>
        </Navbar.Brand>
        <Navbar.Collapse className="">
          <Navbar.Link as={Link} to="/" className="text-lg">
            Home
          </Navbar.Link>
          <Navbar.Link as={Link} to="/search" className="text-lg">
            Search
          </Navbar.Link>
          <Navbar.Link>
            <MegaMenu.Dropdown
              toggle={<span className="text-lg">Category</span>}
            >
              <ul className="grid grid-cols-3 text-lg">
                <div className="space-y-4 p-4">
                  <li>
                    <Link to="/category/medical" className="hover:text-gray-700 flex">
                      <BiPlusMedical  className="mt-1 mr-1"/>
                      Medical
                    </Link>
                  </li>
                  <li>
                    <Link to="/category/emergency" className="hover:text-gray-700 flex">
                      <FaTruckMedical className="mt-1 mr-1"/>
                      Emergency
                    </Link>
                  </li>
                  <li>
                    <Link to="/category/education" className="hover:text-gray-700 flex">
                      <FaGraduationCap  className="mt-1 mr-1"/>
                      Education
                    </Link>
                  </li>
                </div>
                <div className="space-y-4 p-4">
                  <li>
                    <Link to="/category/animals" className="hover:text-gray-700 flex">
                      <PiDogFill  className="mt-1 mr-1"/>
                      Animals
                    </Link>
                  </li>
                  <li>
                    <Link to="/category/competition" className="hover:text-gray-700 flex">
                      <MdEmojiEvents className="mt-1 mr-1"/>
                      Competition
                    </Link>
                  </li>
                  <li>
                    <Link to="/category/event" className="hover:text-gray-700 flex">
                      <MdEventAvailable  className="mt-1 mr-1"/>
                      Event
                    </Link>
                  </li>
                </div>
                <div className="space-y-4 p-4">
                  <li>
                    <Link to="/category/environment" className="hover:text-gray-700 flex">
                      <FaTree className="mt-1 mr-1"/>
                      Environment
                    </Link>
                  </li>
                  <li>
                    <Link to="/category/travel" className="hover:text-gray-700 flex">
                      <GiSailboat className="mt-1 mr-1"/>
                      Travel
                    </Link>
                  </li>
                  <li>
                    <Link to="/category/business" className="hover:text-gray-700 flex">
                      <MdBusinessCenter  className="mt-1 mr-1"/>
                      Business
                    </Link>
                  </li>
                </div>
              </ul>
            </MegaMenu.Dropdown>
          </Navbar.Link>
          {/*<Navbar.Link as={Link} to="/about" className="text-lg">
            About
          </Navbar.Link>*/}
        </Navbar.Collapse>
        <div className="flex md:order-2 flex-grow basis-0 justify-end">
          {user ? (
            <>
            <div className="z-20">
              <Notification />

            </div>
              <Dropdown
              className="z-20"
                inline
                label={
                  <Avatar
                    alt="User settings"
                    img={(props) => (
                      <img
                        alt=""
                        height="48"
                        width="48"
                        referrerPolicy="no-referrer"
                        src={
                          user.profile_picture
                            ? user.profile_picture
                            : "/avatar.svg"
                        }
                        {...props}
                      />
                    )}
                    rounded
                  >
                    <div className="text-sm font-medium">{user.first_name}</div>
                  </Avatar>
                }
              >
                <Dropdown.Item>
                  <Link to="/fundraisers">Your Fundraisers</Link>
                </Dropdown.Item>
                <Dropdown.Item>
                  <Link to="/start-fundraiser">Start a Fundraiser</Link>
                </Dropdown.Item>
                {user.verification_status !== "verified" && (
                  <Dropdown.Item>
                    <Link to="/account/verify">Verify your account</Link>
                  </Dropdown.Item>
                )}
                <Dropdown.Item>
                  <Link to="/account/settings">Account Settings</Link>
                </Dropdown.Item>
                <Dropdown.Divider />
                <Dropdown.Item>
                  <Link to="/logout">Log out</Link>
                </Dropdown.Item>
              </Dropdown>
            </>
          ) : (
            <>
              <Navbar.Collapse>
                <Navbar.Link as={Link} to="/login">
                  <div className="mt-4 text-lg">Sign In</div>
                </Navbar.Link>
                <Navbar.Link as={Link} to="/register" className="mt-3">
                  <Button color="dark">
                    Get Started
                  </Button>
                </Navbar.Link>
              </Navbar.Collapse>
            </>
          )}
          <Navbar.Toggle />
        </div>
      </div>
    </MegaMenu>
  );
};

export default Header;
