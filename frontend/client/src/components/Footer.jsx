import { Footer as FBFooter } from "flowbite-react";

const Footer = () => {
  return (
    <FBFooter container className="bg-gray-200">
      <FBFooter.Copyright href="#" by="Donorbox™" year={2024} />
      <FBFooter.LinkGroup>
        <FBFooter.Link href="/about">About</FBFooter.Link>
        <FBFooter.Link href="#">Privacy Policy</FBFooter.Link>
        <FBFooter.Link href="#">Licensing</FBFooter.Link>
        <FBFooter.Link href="#">Contact</FBFooter.Link>
      </FBFooter.LinkGroup>
    </FBFooter>
  );
};

export default Footer;
