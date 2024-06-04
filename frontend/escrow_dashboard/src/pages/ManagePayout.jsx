import { TabGroup, TabList, Tab, TabPanels, TabPanel } from "@tremor/react";
import { Modal } from "flowbite-react";
import { useEffect, useState } from "react";
import PendingPayoutTable from "../components/payout/PendingPayoutTable";
import projectService from "../services/project";
import { useAuthContext } from "../context/AuthContext";

// eslint-disable-next-line no-unused-vars
const data = [
  {
    id: "1b17d262-e7da-4d66-aa9e-77f03d62e581",
    title: "Support John's Cardiac Treatment",
    cover_picture:
      "https://images.pexels.com/photos/1350560/pexels-photo-1350560.jpeg",
    goal_amount: 10000000,
    accumulated_amount: 11000000,
    end_date: "27/04/2024",
    card_id: "ffffff",
    category_id: 1,
    province: "Vinh",
    country: "Vietnam",
    total_backing: 21,
  },
  {
    id: "1b17d262-e7da-4d66-aa9e-77f03d62e582",
    title: "Find a safe home for these stray cats",
    cover_picture:
      "https://images.pexels.com/photos/5519260/pexels-photo-5519260.jpeg",
    goal_amount: 20000000,
    accumulated_amount: 20000000,
    end_date: "27/04/2024",
    category_id: 3,
    province: "Hanoi",
    country: "Vietnam",
    total_backing: 34,
  },
  {
    id: "1b17d262-e7da-4d66-aa9e-77f03d62e583",
    title: "Support our football club's trip to Worlds!",
    cover_picture:
      "https://images.pexels.com/photos/262524/pexels-photo-262524.jpeg",
    goal_amount: 50000000,
    accumulated_amount: 60000000,
    end_date: "27/04/2024",
    category_id: 5,
    province: "Ho Chi Minh",
    country: "Vietnam",
    total_backing: 28,
  },
];

const ManagePayout = () => {
  const { token } = useAuthContext();
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [isFailed, setIsFailed] = useState(false);
  const [pendingProjects, setPendingProjects] = useState();

  useEffect(() => {
    const fetchProject = async () => {
      try {
        const response = await projectService.getEnded();
        setPendingProjects(
          response.projects.filter((p) => p.current_amount >= p.goal_amount)
        );
      } catch (err) {
        console.error(err);
      }
    };

    fetchProject();
  }, []);

  return (
    <div className="p-10 bg-slate-200 w-full space-y-10 font-sans min-h-screen">
      <div className="text-3xl font-semibold tracking-tight">
        Payout Management
      </div>
      <div className="bg-slate-50 px-5 py-2">
        <TabGroup>
          <TabList className="mt-4" variant="line">
            <Tab>Pending Payouts</Tab>
            <Tab>Disputed Payouts</Tab>
          </TabList>
          <TabPanels>
            <TabPanel>
              <PendingPayoutTable
                token={token}
                data={pendingProjects}
                setIsSuccessful={setIsSuccessful}
                setIsFailed={setIsFailed}
              />
            </TabPanel>
            <TabPanel className="text-sm">To be implemented...</TabPanel>
          </TabPanels>
        </TabGroup>
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

export default ManagePayout;
