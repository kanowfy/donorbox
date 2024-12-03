import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Button, FileInput, Modal } from "flowbite-react";
import { useAuthContext } from "../context/AuthContext";
import uploadService from "../services/upload";

const VerifyAccount = () => {
  const navigate = useNavigate();
  const { user, token } = useAuthContext();
  const [document, setDocument] = useState();
  const [isLoading, setIsLoading] = useState(false);
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [isFailed, setIsFailed] = useState(false);

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
      setIsLoading(false);
      setIsSuccessful(true);
      console.log(response);
      setTimeout(() => {
        navigate("/");
      }, 5000);
    } catch(err) {
      setIsLoading(false);
      setIsFailed(true);
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
                To start your own fundraising campaigns, we need to verify your identity.
                Please print and fill out the attached document and upload it using this form so our team
                can begin the verification process for your account.
              </div>
              <div>
                Click{" "}
                <a
                  className="text-blue-700 hover:underline"
                  href="/verification_form.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
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
      <Modal
        show={isSuccessful}
        size="md"
        onClose={() => {
          setIsSuccessful(false);
          navigate("/");
        }}
        popup
      >
        <Modal.Header />
        <Modal.Body>
          <div className="text-center flex flex-col space-y-2">
            <img
              src="/success.svg"
              height={32}
              width={32}
              className="mx-auto"
            />
            <h3 className="mb-5 text-xl font-normal text-gray-500 dark:text-gray-400">
              Document submitted!
            </h3>
            <p className="text-sm text-gray-600">
              Please wait while your verification is reviewed, we will notify you as soon as possible.
              You can continue to use other functionalities
            </p>
          </div>
        </Modal.Body>
      </Modal>
      <Modal show={isFailed} size="md" onClose={() => setIsFailed(false)} popup>
        <Modal.Header />
        <Modal.Body>
          <div className="text-center flex flex-col space-y-2">
            <img src="/fail.svg" height={32} width={32} className="mx-auto" />
            <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
              Failed to submit document. Please try again later!
            </h3>
          </div>
        </Modal.Body>
      </Modal>
    </div>
  );
};

export default VerifyAccount;
