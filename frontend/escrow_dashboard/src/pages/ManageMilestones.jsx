import { Modal, Tabs } from "flowbite-react";
import { useEffect, useState } from "react";
import projectService from "../services/project";
import { useAuthContext } from "../context/AuthContext";
import MilestoneTable from "../components/MilestoneTable";
import { HiMiniCurrencyDollar } from "react-icons/hi2";
import { HiDocumentSearch } from "react-icons/hi";
import { Card } from "@tremor/react";
import MilestoneProofTable from "../components/MilestoneProofTable";

const ManageMilestones = () => {
  const { token } = useAuthContext();
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [isFailed, setIsFailed] = useState(false);
  const [milestones, setMilestones] = useState();
  const [proofedMilestones, setProofedMilestones] = useState();

  const hasPendingProof = (m) => {
    if (m.milestone.status != "fund_released") {
      return false;
    }

    if (!m.milestone.spending_proofs) {
      return false;
    }

    if (m.milestone.spending_proofs.filter(p => p.status === "pending").length === 0) {
      return false;
    }

    return true;
  }

  useEffect(() => {
    const fetchMilestones = async () => {
      try {
        const response = await projectService.getFundedMilestones(token);
        setMilestones(response.milestones);

        setProofedMilestones(response.milestones.filter(m => hasPendingProof(m)));
        console.log(response.milestones.filter(m => hasPendingProof(m)));
      } catch (err) {
        console.error(err);
      }
    };

    fetchMilestones();
  }, []);

  return (
    <div className="p-10 bg-slate-200 w-full space-y-10 font-sans min-h-screen">
      <div className="text-3xl font-semibold tracking-tight">
        Pending Milestones
      </div>
      <div className="px-5">
      <Card className="shadow-lg shadow-gray-500">
        <Tabs variant="underline">
          <Tabs.Item active title="Resolve Funds" icon={HiMiniCurrencyDollar}>
            <MilestoneTable
              token={token}
              data={milestones?.filter(m => m.milestone.status === "pending")}
              setIsSuccessful={setIsSuccessful}
              setIsFailed={setIsFailed}
            />
          </Tabs.Item>
          <Tabs.Item title="Resolve Proofs" icon={HiDocumentSearch}>
            <MilestoneProofTable
              token={token}
              data={proofedMilestones}
              setIsSuccessful={setIsSuccessful}
              setIsFailed={setIsFailed}
            />
          </Tabs.Item>
        </Tabs>
        </Card>
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

export default ManageMilestones;
