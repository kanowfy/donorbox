import { MegaMenu, Dropdown, Navbar, Avatar, Button } from "flowbite-react";
import { useAuthContext } from "../context/AuthContext";
import { Link } from "react-router-dom";

const Header = () => {
  const { user } = useAuthContext();

  return (
    <MegaMenu className="my-3 mx-16">
      <div className="w-full flex items-center justify-between py-4 px-10 md:space-x-8 shadow-gray-400 shadow-lg rounded-lg">
        <Navbar.Brand href="/">
          <img
            src="/logo.svg"
            className="mr-3 h-6 sm:h-9"
            alt="Flowbite React Logo"
          />
          <span className="self-center whitespace-nowrap text-xl font-semibold">
            Donorbox
          </span>
        </Navbar.Brand>
        <Navbar.Collapse>
          <Navbar.Link href="/" active className="flex flex-col justify-end">
            Home
          </Navbar.Link>
          <Navbar.Link href="/search">Search</Navbar.Link>
          <Navbar.Link>
            <MegaMenu.Dropdown toggle={<>Category</>}>
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
          <Navbar.Link href="/about">About</Navbar.Link>
        </Navbar.Collapse>
        <div className="flex md:order-2">
          {user ? (
            <>
              <Dropdown
                inline
                label={
                  <Avatar
                    alt="User settings"
                    img="https://flowbite.com/docs/images/people/profile-picture-5.jpg"
                    rounded
                  >
                    <div className="text-sm font-medium">Forsen</div>
                  </Avatar>
                }
              >
                <Dropdown.Item>
                  <Link to="#">Your Fundraisers</Link>
                </Dropdown.Item>
                <Dropdown.Item>
                  <Link to="#">Start a Fundraiser</Link>
                </Dropdown.Item>
                <Dropdown.Item>
                  <Link to="#">Account Settings</Link>
                </Dropdown.Item>
                <Dropdown.Divider />
                <Dropdown.Item>
                  <Link to="#">Sign out</Link>
                </Dropdown.Item>
              </Dropdown>
            </>
          ) : (
            <>
              <Navbar.Collapse>
                <Navbar.Link href="/login">
                  <div className="mt-3">Sign In</div>
                </Navbar.Link>
                <Navbar.Link href="#">
                  <Button outline gradientDuoTone="greenToBlue" pill>
                    Start a Fundraiser
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
