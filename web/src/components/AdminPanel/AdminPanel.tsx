import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { MapService } from "../../api";
import AirportGraph from "../AirportGraph/AirportGraph";
import TransportControlPanel from "../TransportControlPanel/TransportControlPanel";
import { Button } from "@/components/ui/button";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { PacmanLoader } from "react-spinners";

const AdminPanel = () => {
  const { data: map, refetch: refetchMap } = useQuery({
    queryKey: ["airportMap"],
    queryFn: MapService.mapGetAirportMap,
  });

  const [activeTab, setActiveTab] = useState("graph");

  return (
    <div className="p-4">
      {map ? (
        <Tabs value={activeTab} onValueChange={setActiveTab}>
          <TabsList className="flex justify-between mb-4 gap-4">
            <div className="flex space-x-4">
              <TabsTrigger value="graph">Airport Graph</TabsTrigger>
              <TabsTrigger value="control">Transport Control</TabsTrigger>
            </div>
            <Button onClick={() => refetchMap()}>Refetch</Button>
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
