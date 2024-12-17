import { Modal } from "flowbite-react";
import { useEffect, useState } from "react";
import projectService from "../services/project";
import { useAuthContext } from "../context/AuthContext";
import DisputeTable from "../components/DisputeTable";

const ManageDisputes = () => {
  const { token } = useAuthContext();
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [isFailed, setIsFailed] = useState(false);
  const [projects, setProjects] = useState();

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const response = await projectService.getDisputed(token);
        setProjects(response.projects);
      } catch (err) {
        console.error(err);
      }
    };

    fetchProjects();
  }, []);

  return (
    <div className="p-10 bg-slate-200 w-full space-y-10 font-sans min-h-screen">
      <div className="text-3xl font-semibold tracking-tight">
        Disputed Fundraisers
      </div>
      <div className="px-5">
        <DisputeTable
          token={token}
          data={projects}
          setIsSuccessful={setIsSuccessful}
          setIsFailed={setIsFailed}
        />
        <Modal
          show={isSuccessful}
          size="md"
          onClose={() => setIsSuccessful(false)}
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
              <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
                Action completed
              </h3>
            </div>
          </Modal.Body>
        </Modal>
        <Modal
          show={isFailed}
          size="md"
          onClose={() => setIsFailed(false)}
          popup
        >
          <Modal.Header />
          <Modal.Body>
            <div className="text-center flex flex-col space-y-2">
              <img src="/fail.svg" height={32} width={32} className="mx-auto" />
              <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
                Failed to complete action
              </h3>
            </div>
          </Modal.Body>
        </Modal>
      </div>
    </div>
  );
};

export default ManageDisputes;
