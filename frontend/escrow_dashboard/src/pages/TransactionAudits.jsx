import { Badge, Table } from "flowbite-react";
import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import transactionService from "../services/transaction";
import utils from "../utils/utils";
import { Button, DateRangePicker } from "@tremor/react";
import { LuRefreshCw } from "react-icons/lu";

const statusFormat = {
  backing: {
    name: "Backing",
    color: "info",
  },
  payout: {
    name: "Payout",
    color: "success",
  },
  refund: {
    name: "Refund",
    color: "pink",
  },
};

const TransactionAudits = () => {
  const { token } = useAuthContext();
  const [rangedTransactions, setRangedTransactions] = useState();
  const [transactions, setTransactions] = useState();
  const [dateRange, setDateRange] = useState();

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const response = await transactionService.getAll(token);
        console.log(response.transactions);
        setTransactions(response.transactions);
        setRangedTransactions(response.transactions);
      } catch (err) {
        console.error(err);
      }
    };

    fetchTransactions();
  }, [token]);

  const handleDateRange = (dateRange) => {
    setDateRange(dateRange);
    if (!dateRange) {
      return;
    }
    const from = dateRange.from;
    let pickedTo = dateRange?.to;
    let to;
    if (!pickedTo) {
      to = new Date(from);
    } else {
      to = new Date(pickedTo);
    }
    to.setDate(to.getDate() + 1);

    console.log(to);

    setRangedTransactions(
      transactions.filter((t) => {
        const tdate = new Date(utils.parseDateFromRFC3339(t.created_at));
        return tdate > from && tdate < to;
      })
    );
  };

  return (
    <div className="p-10 bg-slate-200 w-full font-sans min-h-screen">
      <div className="text-3xl font-semibold tracking-tight mb-10">
        Transaction Logs
      </div>
      <div className="flex justify-center space-x-2 mb-2">
        <DateRangePicker value={dateRange} onValueChange={handleDateRange} />
        <Button
          icon={LuRefreshCw}
          onClick={() => {
            setDateRange(null);
            setRangedTransactions(transactions);
          }}
        >
          Reset
        </Button>
      </div>
      <div className="overflow-x-auto">
        <Table hoverable>
          <Table.Head className="w-fit">
            <Table.HeadCell>Date</Table.HeadCell>
            <Table.HeadCell>Type</Table.HeadCell>
            <Table.HeadCell>Amount</Table.HeadCell>
            <Table.HeadCell>Initiator Card ID</Table.HeadCell>
            <Table.HeadCell>Recipient Card ID</Table.HeadCell>
            <Table.HeadCell>
              <span className="sr-only">View Details</span>
            </Table.HeadCell>
          </Table.Head>
          <Table.Body className="divide-y">
            {rangedTransactions?.map((t) => (
              <Table.Row
                key={t.id}
                className="bg-white dark:border-gray-700 dark:bg-gray-800"
              >
                <Table.Cell className="whitespace-nowrap font-medium text-gray-900 dark:text-white">
                  {utils.formatDateTime(
                    new Date(utils.parseDateFromRFC3339(t.created_at))
                  )}
                </Table.Cell>
                <Table.Cell>
                  <Badge
                    className="w-fit"
                    color={statusFormat[t.transaction_type].color}
                  >
                    {statusFormat[t.transaction_type].name}
                  </Badge>
                </Table.Cell>
                <Table.Cell className="text-black">
                  ₫{utils.formatNumber(t.amount)}
                </Table.Cell>
                <Table.Cell>{t.initiator_card_id.slice(0, 14)}...</Table.Cell>
                <Table.Cell>{t.recipient_card_id.slice(0, 14)}...</Table.Cell>
                <Table.Cell>
                  <button className="font-medium text-blue-600 hover:underline">
                    View
                  </button>
                </Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </div>
    </div>
  );
};

export default TransactionAudits;
