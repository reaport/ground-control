import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { MapService } from "../../api";
import AirportGraph from "../AirportGraph/AirportGraph";
import TransportControlPanel from "../TransportControlPanel/TransportControlPanel";
import { Button } from "@/components/ui/button";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { PacmanLoader } from "react-spinners";

const AdminPanel = () => {
  const [activeTab, setActiveTab] = useState("graph");
  const [refreshInterval, setRefreshInterval] = useState<number | null>(1000);
  const [autoRefresh, setAutoRefresh] = useState(false);

  const { data: map, refetch: refetchMap } = useQuery({
    queryKey: ["airportMap"],
    queryFn: MapService.mapGetAirportMap,
    refetchInterval: autoRefresh ? refreshInterval || 5000 : false,
  });

  return (
    <div className="p-4">
      {map ? (
        <Tabs value={activeTab} onValueChange={setActiveTab}>
          <TabsList className="flex justify-between mb-4 gap-4">
            <div className="flex space-x-4">
              <TabsTrigger value="graph">Airport Graph</TabsTrigger>
              <TabsTrigger value="control">Transport Control</TabsTrigger>
            </div>
            <div className="flex space-x-4 items-center">
              <Button onClick={() => refetchMap()}>Refetch</Button>
              <label className="flex items-center space-x-2">
                <input
                  type="checkbox"
                  checked={autoRefresh}
                  onChange={() => setAutoRefresh(!autoRefresh)}
                />
                <span>Auto-refresh</span>
              </label>
              {autoRefresh && (
                <input
                  type="number"
                  className="border rounded px-2 py-1 w-24"
                  value={refreshInterval || ""}
                  onChange={(e) =>
                    setRefreshInterval(
                      e.target.value ? Number(e.target.value) : null
                    )
                  }
                  min={1000}
                  step={1000}
                  placeholder="Interval"
                  disabled={!autoRefresh}
                />
              )}
            </div>
          </TabsList>

          <TabsContent value="graph">
            <AirportGraph data={map} />
          </TabsContent>
          <TabsContent value="control">
            <TransportControlPanel airportMap={map} />
          </TabsContent>
        </Tabs>
      ) : (
        <PacmanLoader />
      )}
    </div>
  );
};

export default AdminPanel;
