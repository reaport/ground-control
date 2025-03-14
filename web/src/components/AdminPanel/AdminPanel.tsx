import { useQuery } from "@tanstack/react-query";
import { MapService } from "../../api";
// import AirportGraph from "../AirportGraph/AirportGraph";
import TransportControlPanel from "../TransportControlPanel/TransportControlPanel";
import { Button } from "@/components/ui/button";
import { Toaster } from "sonner";

const AdminPanel = () => {
  const { data: map, refetch: refetchMap } = useQuery({
    queryKey: ["adminPanel"],
    queryFn: MapService.mapGetAirportMap,
  });

  return (
    <pre>
      <Button onClick={() => refetchMap()}>Refetch</Button>
      {/* {map! && <AirportGraph data={map} />} */}
      {map! && <TransportControlPanel airportMap={map}/>}
      <Toaster />
    </pre>
  );
};

export default AdminPanel;
