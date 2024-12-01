import { Button, FileInput } from "flowbite-react";
import { useAuthContext } from "../context/AuthContext";

const VerifyAccount = () => {
  const { user } = useAuthContext();
  return (
    <div className="bg-gray-100 h-screen">
      <div className="flex flex-col items-center space-y-10 py-10">
        <div className=" max-w-3xl space-y-5 bg-white p-6 border text-lg">
          <div className="flex justify-center mb-10">
            <div className="text-4xl font-bold underline">Account Verification</div>
          </div>

          {!user ?
          <>
          <div>
            Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nam quidem
            voluptatibus voluptatem voluptate incidunt veniam commodi culpa,
            voluptatum provident fugiat reiciendis itaque nostrum porro laborum
            rerum officiis quae? Debitis, reprehenderit?
          </div>
          <div>
            Click{" "}
            <a
              className="text-blue-700 hover:underline"
              href="www.google.com"
              target="_blank"
            >
              here
            </a>{" "}
            to download the document.
          </div>
          <form>
            <div>Upload your verification document here:</div>
            <FileInput id="upload" helperText="DOC, DOCX or PDF." />
            <Button
              type="submit"
              className="mx-auto my-3"
              color="light"
              pill
              size="xl"
            >
              Submit document
            </Button>
          </form>
          </>
        : <div className="text-xl mt-10">Your account verification is in progress, please be patient.</div>}
        </div>
      </div>
    </div>
  );
};

export default VerifyAccount;
