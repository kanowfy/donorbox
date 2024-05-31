import { DonutChart, Legend } from "@tremor/react";

const project_status = [
  {
    name: "Ongoing",
    quantity: 100,
  },
  {
    name: "Completed Payout",
    quantity: 50,
  },
  {
    name: "Completed Refund",
    quantity: 20,
  },
];

const DashboardPieChart = () => {
  return (
    <div className="bg-white rounded-xl py-10 pl-10 w-fit space-y-5">
      <div className="text-base font-semibold">Project Status Distribution</div>
      <div className="flex items-center justify-center space-x-6">
        <DonutChart
          data={project_status}
          category="quantity"
          index="name"
          colors={["blue", "indigo", "fuchsia"]}
          className="w-60 h-60"
          variant="pie"
        />
        <Legend
          categories={["Ongoing", "Completed Payout", "Completed Refund"]}
          colors={["blue", "indigo", "fuchsia"]}
          className="max-w-xs"
        />
      </div>
    </div>
  );
};

export default DashboardPieChart;
