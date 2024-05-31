import { BarChart } from "@tremor/react";

const chartdata = [
  {
    name: "Disbursed",
    "Number of Fundraisers": 50,
  },
  {
    name: "Refunded",
    "Number of Fundraisers": 20,
  },
];

const DashboardBarChart = () => {
  return (
    <div className="bg-white rounded-xl p-10 w-full space-y-5">
      <h3 className="text-lg font-medium text-tremor-content-strong dark:text-dark-tremor-content-strong">
        Number of completed Fundraisers
      </h3>
      <BarChart
        className="h-60"
        data={chartdata}
        index="name"
        categories={["Number of Fundraisers"]}
        colors={["blue"]}
        yAxisWidth={48}
      />
    </div>
  );
};

export default DashboardBarChart;
