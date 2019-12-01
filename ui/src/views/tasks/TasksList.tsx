import * as React from "react";
import { useEffect, useState } from "react";

import { Table, Pane } from "evergreen-ui";

export const TasksList: React.FunctionComponent = () => {
  const [tasks, setTasks] = useState([]);
  useEffect(() => {
    window.renderer.on("response.download_list", (evt, val) => {
      setTasks(val);
    });

    window.renderer.send({
      evt: "request.download_list",
      val: JSON.stringify({
        type: "url"
      })
    });
  });

  return (
    <Pane flex="1">
      <Table>
        <Table.Head>
          <Table.TextHeaderCell>ID</Table.TextHeaderCell>
          <Table.TextHeaderCell>Task Type</Table.TextHeaderCell>
          <Table.TextHeaderCell>Done?</Table.TextHeaderCell>
        </Table.Head>
        <Table.Body height={500}>
          {tasks.map(task => (
            <Table.Row
              key={task.ID}
              isSelectable
              onSelect={() => alert(task.ID)}
            >
              <Table.TextCell>{task.ID}</Table.TextCell>
              <Table.TextCell>{task.TaskType}</Table.TextCell>
              <Table.TextCell>{task.IsDone ? "true" : "false"}</Table.TextCell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </Pane>
  );
};
