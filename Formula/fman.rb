class Fman < Formula
  desc "Perform fuzzy search on man pages."
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/fman-83d14eee6d58c58f5f7fac20ce5be6484e31918b/fman",
  version "83d14eee6d58c58f5f7fac20ce5be6484e31918b"

  depends_on "apropos"
  depends_on "fzf"

  def install
    bin.install "tools/fman/fman.sh" => "fman"
  end
end
