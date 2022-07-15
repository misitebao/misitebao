import webbrowser

path = "https://misitebao.com"


def main():
    try:
        webbrowser.open(path, new=0, autoraise=True)
    except ZeroDivisionError as err:
        print("An error occurred when opening '%s': %s" % (path, err))
    else:
        print("Opened '%s' successfully." % path)


main()
