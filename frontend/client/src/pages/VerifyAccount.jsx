import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Button, FileInput } from "flowbite-react";
import { useAuthContext } from "../context/AuthContext";
import uploadService from "../services/upload";

const VerifyAccount = () => {
  const navigate = useNavigate();
  const { user, token } = useAuthContext();
  const [document, setDocument] = useState();
  const [isLoading, setIsLoading] = useState(false);

  /*
  useEffect(() => {
    if (!user) {
      navigate("/")
    }
  }, []);
  */


  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    try {
      const response = await uploadDocument(document);
      console.log(response);
    } catch(err) {
      console.error(err);
    }
  }

  const uploadDocument = async (document) => {
    if (!document) {
      throw new Error("Missing document");
    }

    console.log(document);

    const formData = new FormData();
    formData.append("file", document);

    const response = await uploadService.uploadDocument(token, formData);
    return response;
  };

  const onSelectDocument = (e) => {
    if (!e.target.files || e.target.files.length == 0) {
      setDocument(undefined);
      return;
    }

    setDocument(e.target.files[0]);
  }

  return (
    <div className="bg-gray-100 h-screen">
      <div className="flex flex-col items-center space-y-10 py-10">
        <div className=" max-w-3xl space-y-5 bg-white p-6 border text-lg">
          <div className="flex justify-center mb-10">
            <div className="text-4xl font-bold underline">
              Account Verification
            </div>
          </div>

          {user?.verification_status == "unverified" ? (
            <>
              <div>
                Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nam
                quidem voluptatibus voluptatem voluptate incidunt veniam commodi
                culpa, voluptatum provident fugiat reiciendis itaque nostrum
                porro laborum rerum officiis quae? Debitis, reprehenderit?
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
              <form onSubmit={handleSubmit}>
                <div>Upload your verification document here:</div>
                <FileInput
                  accept="application/msword, application/vnd.openxmlformats-officedocument.wordprocessingml.document, application/pdf"
                  onChange={onSelectDocument}
                  id="upload"
                  helperText="DOC, DOCX or PDF."
                />
                <Button
                  type="submit"
                  className="mx-auto my-3"
                  color="light"
                  pill
                  size="xl"
                  isProcessing={isLoading}
                >
                  Submit document
                </Button>
              </form>
            </>
          ) : (
            <div className="text-xl mt-10">
              Your account verification is in progress, please be patient.
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default VerifyAccount;
