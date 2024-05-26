import {
  TabGroup,
  TabList,
  Tab,
  TabPanels,
  TabPanel,
  Dialog,
  DialogPanel,
} from "@tremor/react";
import { useState } from "react";
import PendingPayoutTable from "../components/payout/PendingPayoutTable";
import DisputedPayoutTable from "../components/payout/DisputedPayoutTable";

const data = [
  {
    id: "1b17d262-e7da-4d66-aa9e-77f03d62e582",
    title: "Support John's Cardiac Treatment",
    cover_picture:
      "https://images.pexels.com/photos/1350560/pexels-photo-1350560.jpeg",
    goal_amount: 10000000,
    accumulated_amount: 11000000,
    end_date: "27/04/2024",
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
    total_backing: 34,
  },
  {
    id: "1b17d262-e7da-4d66-aa9e-77f03d62e582",
    title: "Support our football club's trip to Worlds!",
    cover_picture:
      "https://images.pexels.com/photos/262524/pexels-photo-262524.jpeg",
    goal_amount: 50000000,
    accumulated_amount: 60000000,
    end_date: "27/04/2024",
    total_backing: 28,
  },
];

const ManagePayout = () => {
  const [success, setSuccess] = useState(false);

  return (
    <div className="p-10 bg-slate-200 w-full space-y-10 font-sans h-screen">
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
              <PendingPayoutTable data={data} setSuccess={setSuccess} />
            </TabPanel>
            <TabPanel>
              <DisputedPayoutTable data={data} setSuccess={setSuccess} />
            </TabPanel>
          </TabPanels>
        </TabGroup>
        <Dialog open={success} onClose={(val) => setSuccess(val)} static>
          <DialogPanel>Action completed</DialogPanel>
        </Dialog>
      </div>
    </div>
  );
};

export default ManagePayout;
