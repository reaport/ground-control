import React, { FC, useMemo } from "react"
import ReactFlow, {
  Background,
  Controls,
  useNodesState,
  useEdgesState,
  MarkerType,
  Node,
  Edge,
} from "reactflow"
import "reactflow/dist/style.css"
import dagre from "dagre"

import type { AirportMap } from "../../api" // Adjust to your real path
import { ShadcnNode } from "./ShadcnNode"      // The custom node

// React Flow node dimensions used for dagre layout
const nodeWidth = 200
const nodeHeight = 100

// Configure dagre to lay out left-to-right (you can do "TB" for top-to-bottom, etc.)
const dagreGraph = new dagre.graphlib.Graph()
dagreGraph.setDefaultEdgeLabel(() => ({}))
dagreGraph.setGraph({ rankdir: "LR" })
dagreGraph.setGraph({
  rankdir: "LR",     // or "TB"
  ranksep: 200,      // vertical spacing
  nodesep: 200,      // horizontal spacing
  edgesep: 50,       // spacing between edges
});


type AirportMapProps = {
  airportMap: AirportMap
}

function getLayoutedElements(airportMap: AirportMap) {
  // reset the graph so repeated calls won't accumulate old edges
  dagreGraph.nodes().forEach((node) => dagreGraph.removeNode(node))

  // Build up dagre nodes
  const reactFlowNodes: Node[] = airportMap.nodes.map((n) => {
    dagreGraph.setNode(n.id, { width: nodeWidth, height: nodeHeight })
    return {
      id: n.id,
      type: "shadcnNode", // so ReactFlow knows to use our custom node component
      data: { 
        label: n.id, 
        types: n.types, 
        vehicles: n.vehicles 
      },
      position: { x: 0, y: 0 }, // will be set by dagre layout
    }
  })

  // Build up dagre edges
  const reactFlowEdges: Edge[] = airportMap.edges.map((edge, index) => {
    // The 'id' of an edge must be unique
    const edgeId = `edge-${index}`

    // Add the edge to dagre
    dagreGraph.setEdge(edge.from, edge.to)

    return {
      id: edgeId,
      source: edge.from,
      target: edge.to,
      label: `${edge.distance}m`,
      type: 'smoothstep',     // or 'straight' or 'step'
      markerEnd: { type: MarkerType.Arrow },
    }    
  })

  // Run dagre layout
  dagre.layout(dagreGraph)

  // Update node positions with dagre's layout results
  reactFlowNodes.forEach((node) => {
    const nodeWithPosition = dagreGraph.node(node.id)
    node.position = {
      x: nodeWithPosition.x - nodeWidth / 2,
      y: nodeWithPosition.y - nodeHeight / 2,
    }
  })

  return { nodes: reactFlowNodes, edges: reactFlowEdges }
}

export const AirportMapView: FC<AirportMapProps> = ({ airportMap }) => {
  // Compute layouted elements only once whenever 'airportMap' changes
  const { nodes: layoutedNodes, edges: layoutedEdges } = useMemo(
    () => getLayoutedElements(airportMap),
    [airportMap]
  )

  // If you want to allow editing or dragging, store them in React Flow’s state:
  const [nodes, setNodes, onNodesChange] = useNodesState(layoutedNodes)
  const [edges, setEdges, onEdgesChange] = useEdgesState(layoutedEdges)

  // Register your custom node types
  const nodeTypes = useMemo(
    () => ({ shadcnNode: ShadcnNode }),
    []
  )

  return (
    <div style={{ width: "100%", height: "100%" }}>
      <ReactFlow
        nodeTypes={nodeTypes}
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        fitView
      >
        <Background />
        <Controls />
      </ReactFlow>
    </div>
  )
}
