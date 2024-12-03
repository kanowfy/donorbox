import { MegaMenu, Dropdown, Navbar, Avatar, Button } from "flowbite-react";
import { useAuthContext } from "../context/AuthContext";
import { Link } from "react-router-dom";
import Notification from "./Notification";

const Header = () => {
  const { user } = useAuthContext();
  console.log(user);

  return (
    <MegaMenu
      className="flex justify-center w-full"
      theme={{
        root: {
          base: "bg-gray-50 dark:border-gray-700 dark:bg-gray-800 sm:px-4",
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
          <Navbar.Link>
            <MegaMenu.Dropdown
              toggle={<span className="text-lg">Category</span>}
            >
              <ul className="grid grid-cols-3">
                <div className="space-y-4 p-4">
                  <li>
                    <a href="#" className="hover:text-primary-600">
                      Medical
                    </a>
                  </li>
                  <li>
                    <a href="#" className="hover:text-primary-600">
                      Emergency
                    </a>
                  </li>
                  <li>
                    <a href="#" className="hover:text-primary-600">
                      Education
                    </a>
                  </li>
                </div>
                <div className="space-y-4 p-4">
                  <li>
                    <a href="#" className="hover:text-primary-600">
                      Animals
                    </a>
                  </li>
                  <li>
                    <a href="#" className="hover:text-primary-600">
                      Competition
                    </a>
                  </li>
                  <li>
                    <a href="#" className="hover:text-primary-600">
                      Event
                    </a>
                  </li>
                </div>
                <div className="space-y-4 p-4">
                  <li>
                    <a href="#" className="hover:text-primary-600">
                      Environment
                    </a>
                  </li>
                  <li>
                    <a href="#" className="hover:text-primary-600">
                      Travel
                    </a>
                  </li>
                  <li>
                    <a href="#" className="hover:text-primary-600">
                      Business
                    </a>
                  </li>
                </div>
              </ul>
            </MegaMenu.Dropdown>
          </Navbar.Link>
          <Navbar.Link as={Link} to="/search" className="text-lg">
            Search
          </Navbar.Link>
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
