import { Badge, Table } from "flowbite-react";
import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import auditService from "../services/audit";
import utils from "../utils/utils";
import { Button, DateRangePicker } from "@tremor/react";
import { LuRefreshCw } from "react-icons/lu";

const operationFormat = {
  CREATE: "info",
  UPDATE: "warning",
  DELETE: "pink"
};

const entityFormat = {
  user: "info",
  project: "success",
  milestone: "warning"
};


const AuditTrails = () => {
  const { token } = useAuthContext();
  const [rangedAudits, setRangedAudits] = useState();
  const [audits, setAudits] = useState();
  const [dateRange, setDateRange] = useState();

  useEffect(() => {
    const fetchAudits = async () => {
      try {
        const response = await auditService.getAuditHistory(token);
        console.log(response.audits);
        setAudits(response.audits);
        setRangedAudits(response.audits);
      } catch (err) {
        console.error(err);
      }
    };

    fetchAudits();
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

    setRangedAudits(
      audits.filter((t) => {
        const tdate = new Date(utils.parseDateFromRFC3339(t.created_at));
        return tdate > from && tdate < to;
      })
    );
  };

  const limitValue = (val) => {
    if (val?.length > 25) {
      return `${val.substring(0,25)}...`;
    }

    return val;
  }

  return (
    <div className="p-10 bg-slate-200 w-full font-sans min-h-screen">
      <div className="text-3xl font-semibold tracking-tight mb-10">
        Audit Trails
      </div>
      <div className="flex justify-center space-x-2 mb-2">
        <DateRangePicker value={dateRange} onValueChange={handleDateRange} />
        <Button
          icon={LuRefreshCw}
          onClick={() => {
            setDateRange(null);
            setRangedAudits(audits);
          }}
        >
          Reset
        </Button>
      </div>
      <div className="overflow-x-auto">
        <Table hoverable>
          <Table.Head className="w-fit">
            <Table.HeadCell>ID</Table.HeadCell>
            <Table.HeadCell>Performer ID</Table.HeadCell>
            <Table.HeadCell>Operation</Table.HeadCell>
            <Table.HeadCell>Entity</Table.HeadCell>
            <Table.HeadCell>Entity ID</Table.HeadCell>
            <Table.HeadCell>Field</Table.HeadCell>
            <Table.HeadCell>Old Value</Table.HeadCell>
            <Table.HeadCell>New Value</Table.HeadCell>
            <Table.HeadCell>Time</Table.HeadCell>
            {/*<Table.HeadCell>
              <span className="sr-only">View Details</span>
            </Table.HeadCell>*/}
          </Table.Head>
          <Table.Body className="divide-y">
            {rangedAudits?.map((t) => (
              <Table.Row
                key={t.id}
                className="bg-white dark:border-gray-700 dark:bg-gray-800"
              >
                <Table.Cell>{t.id}</Table.Cell>
                <Table.Cell>{`${t.user_id ? `User (${t.user_id})` : `Escrow (${t.escrow_id})`}`}</Table.Cell>
                <Table.Cell>
                  <Badge
                    className="w-fit"
                    color={operationFormat[t.operation_type]}
                  >
                    {t.operation_type}
                  </Badge>
                </Table.Cell>
                <Table.Cell>
                  <Badge
                    className="w-fit"
                    color={entityFormat[t.entity_type]}
                  >
                  {t.entity_type}
                  </Badge>
                </Table.Cell>
                <Table.Cell>
                  {t?.entity_id}
                </Table.Cell>

                <Table.Cell>
                  {t.field_name && (
                  <Badge
                    className="w-fit"
                    color="success"
                  >
                    {t.field_name}
                  </Badge>
                  )}
                </Table.Cell>

                <Table.Cell>{t.old_value && limitValue(t.old_value)}</Table.Cell>
                <Table.Cell>{t.new_value && limitValue(t.new_value)}</Table.Cell>

                <Table.Cell className="whitespace-nowrap font-medium text-gray-900 dark:text-white">
                  {utils.formatDateTime(
                    new Date(utils.parseDateFromRFC3339(t.created_at))
                  )}
                </Table.Cell>
                {/*<Table.Cell>
                  <button className="font-medium text-blue-600 hover:underline">
                    View
                  </button>
                </Table.Cell>*/}
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </div>
    </div>
  );
};

export default AuditTrails;
