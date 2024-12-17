import { useEffect, useState } from "react";
import { INFURA_APIKEY, CONTRACT_ADDRESS, ABI } from "../constants";
import { ethers } from "ethers";
import { Modal, Table } from "flowbite-react";

const provider = new ethers.InfuraProvider("sepolia", INFURA_APIKEY);

const TransactionDisplay = () => {
  const [milestoneReleaseEvents, setMilestoneReleaseEvents] = useState();
  const [verifiedProofEvents, setVerifiedProofEvents] = useState();
  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const events = await queryChain();
        setMilestoneReleaseEvents(events[0]);
        setVerifiedProofEvents(events[1]);
        console.log("events", events);
      } catch (err) {
        console.error(err);
      }
    };

    fetchEvents();
  }, []);

  return (
    <div className="px-10 pt-10 pb-20 space-y-20 border-y">
      {milestoneReleaseEvents && (
        <>
          <div className="flex flex-col items-center space-y-5">
            <div className="text-cyan-600 text-2xl">Milestone Release Events</div>
            <MilestoneReleaseTable events={milestoneReleaseEvents} />
          </div>
          <div className="flex flex-col items-center space-y-5">
            <div className=" text-lime-600 text-2xl">Verified Proof of Expenditure Events</div>
            <VerifiedProofTable events={verifiedProofEvents} />
          </div>
        </>
      )}
    </div>
  );
};

const MilestoneReleaseTable = ({ events }) => {
  const [showReceipt, setShowReceipt] = useState(false);
  const [receipt, setReceipt] = useState();
  return (
    <Table striped hoverable className="shadow-none text-lg">
      <Table.Head className="w-fit">
        <Table.HeadCell>Transaction Hash</Table.HeadCell>
        <Table.HeadCell>Project</Table.HeadCell>
        <Table.HeadCell>Transfer Receipt</Table.HeadCell>
        <Table.HeadCell>Transfer Note</Table.HeadCell>
        <Table.HeadCell>Date</Table.HeadCell>
      </Table.Head>
      <Table.Body className="divide-y">
        {events.map((e) => (
          <Table.Row key={e.args[0]}>
            <Table.Cell>
              <a
                href={`https://sepolia.etherscan.io/tx/${e.transactionHash}`}
                target="_blank"
                className="hover:text-blue-600 text-gray-800"
              >
                {e.transactionHash.substring(0, 20)}...
              </a>
            </Table.Cell>
            <Table.Cell>
              <a
                href={`http://localhost:4001/fundraiser/${e.args[1]}`}
                target="_blank"
                className="hover:text-blue-600"
              >
                {Number(e.args[1])}
              </a>
            </Table.Cell>
            <Table.Cell>
              <div
                className="hover:text-blue-600 hover:underline cursor-pointer"
                onClick={() => {
                  setShowReceipt(true);
                  setReceipt(e.args[3]);
                }}
              >
                View
              </div>
            </Table.Cell>
            <Table.Cell>{e.args[4]}</Table.Cell>
            <Table.Cell>{e.args[5]}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
      <Modal show={showReceipt} onClose={() => setShowReceipt(false)}>
        <Modal.Header>Transfer Receipt</Modal.Header>
        <Modal.Body>
          <div className="flex justify-center">
            <img src={receipt} alt="receipt" />

          </div>
        </Modal.Body>
      </Modal>
    </Table>
  );
};

const VerifiedProofTable = ({ events }) => {
  const [showReceipt, setShowReceipt] = useState(false);
  const [receipt, setReceipt] = useState();
  return (
    <Table striped hoverable className="shadow-none text-lg">
      <Table.Head className="w-fit">
        <Table.HeadCell>Transaction Hash</Table.HeadCell>
        <Table.HeadCell>Project</Table.HeadCell>
        <Table.HeadCell>Transfer Receipt</Table.HeadCell>
        <Table.HeadCell>Proof Image</Table.HeadCell>
        <Table.HeadCell>Date</Table.HeadCell>
      </Table.Head>
      <Table.Body className="divide-y">
        {events.map((e) => (
          <Table.Row key={e.args[0]}>
            <Table.Cell>
              <a
                href={`https://sepolia.etherscan.io/tx/${e.transactionHash}`}
                target="_blank"
                className="hover:text-blue-600 text-gray-800"
              >
                {e.transactionHash.substring(0, 25)}...
              </a>
            </Table.Cell>
            <Table.Cell>
              <a
                href={`http://localhost:4001/fundraiser/${e.args[1]}`}
                target="_blank"
                className="hover:text-blue-600"
              >
                {Number(e.args[1])}
              </a>
            </Table.Cell>
            <Table.Cell>  
              <div
                className="hover:text-blue-600 hover:underline cursor-pointer"
                onClick={() => {
                  setShowReceipt(true);
                  setReceipt(e.args[3]);
                }}
              >
                View
              </div>
            </Table.Cell>
            <Table.Cell>  
              <div
                className="hover:text-blue-600 hover:underline cursor-pointer"
                onClick={() => {
                  setShowReceipt(true);
                  setReceipt(e.args[4]);
                }}
              >
                View
              </div>
              </Table.Cell>
            <Table.Cell>{e.args[5]}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
      <Modal show={showReceipt} onClose={() => setShowReceipt(false)}>
        <Modal.Header/>
        <Modal.Body>
          <div className="flex justify-center">
            <img src={receipt} alt="receipt" />

          </div>
        </Modal.Body>
      </Modal>
    </Table>
  );
};


const queryChain = async () => {
  const contract = new ethers.Contract(CONTRACT_ADDRESS, ABI, provider);
  const mrFilter = contract.filters.MilestoneReleaseStored();
  const mrEvents = await contract.queryFilter(mrFilter, 0, "latest");

  const vpFilter = contract.filters.VerifiedProofStored();
  const vpEvents = await contract.queryFilter(vpFilter, 0, "latest");

  return [mrEvents, vpEvents];
};

export default TransactionDisplay;
