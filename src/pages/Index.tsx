import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Card } from "@/components/ui/card";
import { useToast } from "@/hooks/use-toast";

const Index = () => {
  const { toast } = useToast();
  const [emailData, setEmailData] = useState({
    source: "",
    destination: "",
    subject: "",
    body: "",
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    try {
      console.log('Attempting to send email with data:', {
        source: emailData.source,
        destination: [emailData.destination],
        message: {
          subject: emailData.subject,
          body: emailData.body,
        },
      });

      const response = await fetch("http://localhost:3000/v1/email/send", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json",
        },
        body: JSON.stringify({
          source: emailData.source,
          destination: [emailData.destination],
          message: {
            subject: emailData.subject,
            body: emailData.body,
          },
        }),
      });

      const data = await response.json();
      console.log('Server response:', data);

      if (!response.ok) {
        if (response.status === 404) {
          throw new Error("API endpoint not found. Please check if the server is running on the correct port.");
        } else if (response.status === 400) {
          throw new Error("Invalid email data. Please check your input.");
        } else {
          throw new Error(`Server error: ${data.error || 'Unknown error'}`);
        }
      }

      toast({
        title: "Success",
        description: "Email sent successfully!"
      });
      
      setEmailData({
        source: "",
        destination: "",
        subject: "",
        body: "",
      });
    } catch (error) {
      console.error("Error sending email:", error);
      if (error instanceof Error) {
        if (error.message.includes("Failed to fetch")) {
          toast({
            variant: "destructive",
            title: "Error",
            description: "Cannot connect to the server. Please make sure the backend server is running (go run cmd/main.go)"
          });
        } else {
          toast({
            variant: "destructive",
            title: "Error",
            description: error.message
          });
        }
      } else {
        toast({
          variant: "destructive",
          title: "Error",
          description: "An unexpected error occurred. Please try again."
        });
      }
    }
  };

  return (
    <div className="min-h-screen p-8 bg-gray-50">
      <div className="max-w-4xl mx-auto space-y-8">
        <h1 className="text-3xl font-bold text-center">AWS SES Mock API Testing</h1>
        
        {/* Email Sending Form */}
        <Card className="p-6">
          <h2 className="text-xl font-semibold mb-4">Send Test Email</h2>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block text-sm font-medium mb-1">From:</label>
              <Input
                type="email"
                value={emailData.source}
                onChange={(e) => setEmailData({ ...emailData, source: e.target.value })}
                placeholder="sender@example.com"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">To:</label>
              <Input
                type="email"
                value={emailData.destination}
                onChange={(e) => setEmailData({ ...emailData, destination: e.target.value })}
                placeholder="recipient@example.com"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Subject:</label>
              <Input
                type="text"
                value={emailData.subject}
                onChange={(e) => setEmailData({ ...emailData, subject: e.target.value })}
                placeholder="Email Subject"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Message:</label>
              <Textarea
                value={emailData.body}
                onChange={(e) => setEmailData({ ...emailData, body: e.target.value })}
                placeholder="Type your message here..."
                required
                rows={4}
              />
            </div>
            <Button type="submit" className="w-full">
              Send Email
            </Button>
          </form>
        </Card>

        {/* Statistics Section */}
        <Card className="p-6">
          <h2 className="text-xl font-semibold mb-4">Email Statistics</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div className="p-4 bg-blue-50 rounded-lg">
              <p className="text-sm text-blue-600">Sent Last 24 Hours</p>
              <p className="text-2xl font-bold">0</p>
            </div>
            <div className="p-4 bg-green-50 rounded-lg">
              <p className="text-sm text-green-600">Delivery Rate</p>
              <p className="text-2xl font-bold">100%</p>
            </div>
            <div className="p-4 bg-purple-50 rounded-lg">
              <p className="text-sm text-purple-600">Quota Remaining</p>
              <p className="text-2xl font-bold">200</p>
            </div>
          </div>
        </Card>
      </div>
    </div>
  );
};

export default Index;