using System.Diagnostics;

namespace main
{
  class Program
  {
    static void Main(string[] args)
    {
      String path = "https://misitebao.com";

      Process p = new Process();

      try
      {
        p.StartInfo.UseShellExecute = true;
        p.StartInfo.FileName = path;
        p.Start();
      }
      catch (Exception e)
      {
        throw new Exception(string.Format("An error occurred when opening '{0}': {1}", path, e));
      }

      Console.WriteLine(string.Format("Opened '{0}' successfully.", path));
    }
  }
}








