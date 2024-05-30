import { Button, Dropdown } from "flowbite-react";
import { useEffect, useState, useReducer } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import projectService from "../../services/project";
import ProjectCard from "../../components/ProjectCard";
import { MdCancel } from "react-icons/md";
import utils from "../../utils/utils";
import { FaSort } from "react-icons/fa";
import { ImSortAlphaAsc, ImSortAlphaDesc } from "react-icons/im";
import sortBy from "sort-by";
import { CategoryIndexMap } from "../../constants";

const Search = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const initialSearchQuery = searchParams.get("q");
  const navigate = useNavigate();
  const [query, setQuery] = useState(initialSearchQuery || "");
  const [projects, setProjects] = useState([]);
  const [filtered, setFiltered] = useState([]);
  const [categoryFilter, setCategoryFilter] = useState(0);

  // eslint-disable-next-line no-unused-vars
  const [_, forceUpdate] = useReducer((x) => x + 1, 0);

  function submitSearchQuery(e) {
    e.preventDefault();
    setSearchParams({ q: query });
    navigate(`/search?q=${encodeURIComponent(query)}`);
  }

  function sortList(sortType) {
    setFiltered(filtered.sort(sortBy(sortType)));
    forceUpdate();
  }

  useEffect(() => {
    const searchProjects = async (q) => {
      try {
        const response = await projectService.search(q, 1, 12);
        setProjects(response.projects);
        setFiltered(response.projects);
      } catch (err) {
        console.log(err);
      }
    };

    if (initialSearchQuery) {
      searchProjects(initialSearchQuery);
    }
  }, [initialSearchQuery]);

  return (
    <>
      <section className="min-h-screen bg-cover bg-gradient-to-b from-white to-sky-200 mt-20">
        <div className="flex flex-col items-center">
          <div className="pb-7 font-semibold text-4xl tracking-tight">
            Search for fundraising projects
          </div>
          <div className="pb-10 font-normal text-lg">
            Find fundraisers by category, description or location
          </div>
          <div>
            <form
              onSubmit={submitSearchQuery}
              className="min-w-128 max-w-md mx-auto"
            >
              <label className="mb-2 text-sm font-medium text-gray-900 sr-only">
                Search
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                  <svg
                    className="w-4 h-4 text-gray-500"
                    aria-hidden="true"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 20 20"
                  >
                    <path
                      stroke="currentColor"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
                    />
                  </svg>
                </div>
                <input
                  type="search"
                  className="block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  placeholder="Search..."
                  value={query}
                  onChange={(e) => setQuery(e.target.value)}
                  required
                />
                <button
                  type="submit"
                  className="text-white absolute end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >
                  Search
                </button>
              </div>
            </form>
          </div>
        </div>
        <section className="mx-32 my-10">
          {initialSearchQuery ? (
            <div>
              <div className="flex gap-1">
                <Dropdown
                  label="Location"
                  dismissOnClick={false}
                  color="light"
                  outline
                  size="sm"
                  pill
                >
                  <Dropdown.Header>
                    <span>Choose location worldwide</span>
                  </Dropdown.Header>
                  <Dropdown.Item>HAHA</Dropdown.Item>
                </Dropdown>
                <Dropdown
                  label="Category"
                  dismissOnClick={false}
                  color="light"
                  outline
                  size="sm"
                  pill
                >
                  <Dropdown.Header>
                    <span>Choose a category</span>
                  </Dropdown.Header>
                  <Dropdown.Item>
                    <div className="grid grid-cols-3 max-w-sm gap-2">
                      {Object.entries(CategoryIndexMap).map(([name, num]) => (
                        <Button
                          size="xs"
                          color="success"
                          key={name}
                          onClick={() => {
                            setCategoryFilter(num);
                            setFiltered(
                              projects.filter((p) => p.category_id == num)
                            );
                          }}
                          pill
                          outline={categoryFilter != num}
                        >
                          {name}
                        </Button>
                      ))}
                    </div>
                  </Dropdown.Item>
                </Dropdown>
                <Button
                  size="sm"
                  color="light"
                  pill
                  onClick={() => {
                    setFiltered(
                      filtered.filter(
                        (p) => p.goal_amount - p.current_amount < 1000000
                      )
                    );
                  }}
                >
                  Close to goal
                </Button>
                <Button
                  size="sm"
                  color="light"
                  pill
                  onClick={() => {
                    setFiltered(
                      filtered.filter((p) =>
                        utils.calculateDayDifference(
                          Date.now(),
                          utils.parseDateFromRFC3339(p.start_date) < 3
                        )
                      )
                    );
                  }}
                >
                  Recently launched
                </Button>
                <Dropdown
                  label={
                    <div className="flex">
                      <FaSort className="mr-2 h-5 w-5" />
                      <div>Sort by</div>
                    </div>
                  }
                  dismissOnClick={false}
                  color="light"
                  outline
                  size="sm"
                  pill
                  arrowIcon={false}
                >
                  <Dropdown.Item
                    onClick={() => {
                      sortList("title");
                    }}
                  >
                    <ImSortAlphaAsc className="mr-2 h-5 w-5" />
                    Alphabetical
                  </Dropdown.Item>
                  <Dropdown.Item
                    onClick={() => {
                      sortList("-title");
                    }}
                  >
                    <ImSortAlphaDesc className="mr-2 h-5 w-5" />
                    Alphabetical
                  </Dropdown.Item>
                  <Dropdown.Item
                    onClick={() => {
                      sortList("-backing_count");
                    }}
                  >
                    <ImSortAlphaAsc className="mr-2 h-5 w-5" />
                    Most backings
                  </Dropdown.Item>
                  <Dropdown.Item
                    onClick={() => {
                      sortList("backing_count");
                    }}
                  >
                    <ImSortAlphaDesc className="mr-2 h-5 w-5" />
                    Least backings
                  </Dropdown.Item>
                  <Dropdown.Item
                    onClick={() => {
                      sortList("-current_amount");
                    }}
                  >
                    <ImSortAlphaAsc className="mr-2 h-5 w-5" />
                    Most amount backed
                  </Dropdown.Item>
                  <Dropdown.Item
                    onClick={() => {
                      sortList("current_amount");
                    }}
                  >
                    <ImSortAlphaDesc className="mr-2 h-5 w-5" />
                    Least amount backed
                  </Dropdown.Item>
                </Dropdown>
                <Button
                  size="sm"
                  color="dark"
                  pill
                  onClick={() => {
                    setCategoryFilter(0);
                    setFiltered(projects);
                  }}
                >
                  <MdCancel className="mr-2 h-5 w-5" />
                  Clear filters
                </Button>
              </div>
              <div className="flex justify-center">
                <div className="grid grid-cols-1 gap-7 md:grid-cols-3 xl:grid-cols-4 mx-16 mt-10 mb-16">
                  {filtered.map((p) => (
                    <ProjectCard
                      id={p.id}
                      title={p.title}
                      cover={p.cover_picture}
                      currentAmount={p.current_amount}
                      goalAmount={p.goal_amount}
                      numBackings={p.backing_count}
                      key={p.id}
                    />
                  ))}
                </div>
              </div>
            </div>
          ) : (
            ""
          )}
        </section>
      </section>
    </>
  );
};

export default Search;
