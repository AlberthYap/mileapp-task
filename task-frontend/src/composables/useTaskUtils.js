export function useTaskUtils() {
  // Utility functions
  const getStatusClass = (status) => {
    if (!status) return "bg-gray-100 text-gray-800";

    const classes = {
      pending: "bg-yellow-100 text-yellow-800",
      in_progress: "bg-blue-100 text-blue-800",
      completed: "bg-green-100 text-green-800",
    };
    return classes[status] || "bg-gray-100 text-gray-800";
  };

  const getStatusDotClass = (status) => {
    if (!status) return "bg-gray-500";

    const classes = {
      pending: "bg-yellow-500",
      in_progress: "bg-blue-500",
      completed: "bg-green-500",
    };
    return classes[status] || "bg-gray-500";
  };

  const getPriorityClass = (priority) => {
    if (!priority) return "bg-gray-100 text-gray-700";

    const classes = {
      low: "bg-gray-100 text-gray-700",
      medium: "bg-orange-100 text-orange-700",
      high: "bg-red-100 text-red-700",
    };
    return classes[priority] || "bg-gray-100 text-gray-700";
  };

  const formatStatus = (status) => {
    if (!status) return "Unknown";

    const labels = {
      pending: "Pending",
      in_progress: "In Progress",
      completed: "Completed",
    };
    return labels[status] || status;
  };

  const formatPriority = (priority) => {
    if (!priority || typeof priority !== "string") return "None";
    return priority.charAt(0).toUpperCase() + priority.slice(1);
  };

  const formatDate = (dateString) => {
    if (!dateString) return "";

    try {
      const date = new Date(dateString);
      if (isNaN(date.getTime())) return "";

      const now = new Date();
      const diffTime = date - now;
      const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));

      if (diffDays < 0) return "Overdue";
      if (diffDays === 0) return "Today";
      if (diffDays === 1) return "Tomorrow";
      if (diffDays <= 7) return `${diffDays} days`;

      return date.toLocaleDateString("en-US", {
        month: "short",
        day: "numeric",
      });
    } catch (error) {
      console.error("Error formatting date:", error);
      return "";
    }
  };

  return {
    getStatusClass,
    getStatusDotClass,
    getPriorityClass,
    formatStatus,
    formatPriority,
    formatDate,
  };
}
