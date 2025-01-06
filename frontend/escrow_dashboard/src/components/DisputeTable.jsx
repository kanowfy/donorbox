import { Card } from "@tremor/react";
import PropTypes from "prop-types";
import escrowService from "../services/escrow";
import { useNavigate } from "react-router-dom";
import { Badge, Table, Dropdown } from "flowbite-react";
import { FaDownload } from "react-icons/fa";
import { IoIosMore } from "react-icons/io";
import { CLIENT_URL } from "../constants";

const ReportStatusMap = {
  pending: {
    color: "warning",
    text: "pending",
  },
  dismissed: {
    color: "failure",
    text: "dismissed",
  },
  resolved: {
    color: "info",
    text: "resolved",
  },
};

const DisputeTable = ({ token, data, setIsSuccessful, setIsFailed }) => {
  const navigate = useNavigate();

  const handleGenerate = async (data) => {
    try {
      const response = await escrowService.generateReport(token, data);
      const url = window.URL.createObjectURL(new Blob([response]));

      // Create a temporary link element
      const link = document.createElement("a");
      link.href = url;
      link.setAttribute("download", "Report.pdf"); // File name for the download
      document.body.appendChild(link);

      // Trigger the download
      link.click();

      // Cleanup
      link.remove();
      window.URL.revokeObjectURL(url);
    } catch (err) {
      console.error(err);
    }
  };

  const handleProjectStopped = async (projectID) => {
    try {
      await escrowService.resolveDispute(token, {
        project_id: Number(projectID),
        mark_stopped: true,
      });
      setIsSuccessful(true);
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      console.error(err);
      setIsFailed(true);
    }
  };

  const handleProjectReconciled = async (projectID) => {
    console.log("dispute project id", projectID)
    try {
      await escrowService.resolveDispute(token, {
        project_id: Number(projectID),
        mark_reconciled: true,
      });
      setIsSuccessful(true);
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      console.error(err);
      setIsFailed(true);
    }
  };

  return (
    <div>
      <Card className="my-5 shadow-lg shadow-gray-500">
        <h3 className="text-tremor-content-strong dark:text-dark-tremor-content-strong">
          List of Fundraisers requiring investigation due to policy violation
        </h3>
        <Table className="mt-5" striped hoverable>
          <Table.Head className="w-fit">
            <Table.HeadCell>ID</Table.HeadCell>
            <Table.HeadCell>Project Title</Table.HeadCell>
            <Table.HeadCell>Owner Name</Table.HeadCell>
            <Table.HeadCell>Dispute Cause</Table.HeadCell>
            <Table.HeadCell>Amount Disputed</Table.HeadCell>
            <Table.HeadCell># Confirmed Reports</Table.HeadCell>
            <Table.HeadCell>Document</Table.HeadCell>
            <Table.HeadCell>Action</Table.HeadCell>
          </Table.Head>
          <Table.Body className="divide-y">
            {data?.map((item) => (
              <Table.Row key={item.project.id}>
                <Table.Cell>{item.project.id}</Table.Cell>
                <Table.Cell>
                  <a
                    href={`${CLIENT_URL}/fundraiser/${item.project.id}`}
                    target="_blank"
                    className="hover:text-blue-600"
                  >
                    {item.project.title}
                  </a>
                </Table.Cell>
                <Table.Cell>{`${item.user?.first_name} ${item.user?.last_name}`}</Table.Cell>
                <Table.Cell>
                  {item.is_reported ? (
                    <Badge className="w-fit" color="failure">
                      Legitimate Report Filed
                    </Badge>
                  ) : (
                    <Badge className="w-fit" color="warning">
                      Failed to Validate Spending
                    </Badge>
                  )}
                </Table.Cell>
                <Table.Cell>₫{item.project.total_fund.toLocaleString()}</Table.Cell>
                <Table.Cell>
                  {item.reports ? item.reports.length : 0}
                </Table.Cell>
                <Table.Cell>
                  <button
                    className="font-semibold hover:underline text-zinc-600 flex items-baseline gap-1"
                    onClick={() => handleGenerate(item)}
                  >
                    Generate Report{" "}
                    <span>
                      <FaDownload className="mt-1" />
                    </span>
                  </button>
                </Table.Cell>
                <Table.Cell>
                  <Dropdown
                    label=""
                    dismissOnClick={false}
                    renderTrigger={() => (
                      <span className="cursor-pointer text-gray-800">
                        <IoIosMore className="w-6 h-6" />
                      </span>
                    )}
                  >
                    <Dropdown.Item>
                      <span
                        className="text-blue-600 font-semibold"
                        onClick={() => handleProjectReconciled(item.project.id)}
                      >
                        Mark as Reconciled
                      </span>
                    </Dropdown.Item>
                    <Dropdown.Item>
                      <span
                        className="text-red-600 font-semibold"
                        onClick={() => handleProjectStopped(item.project.id)}
                      >
                        Mark as Stopped
                      </span>
                    </Dropdown.Item>
                  </Dropdown>
                </Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </Card>
    </div>
  );
};

DisputeTable.propTypes = {
  data: PropTypes.array,
  setIsSuccessful: PropTypes.func,
  setIsFailed: PropTypes.func,
};

export default DisputeTable;
